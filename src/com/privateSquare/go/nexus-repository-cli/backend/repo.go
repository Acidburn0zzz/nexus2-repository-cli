package backend

import (
	m "com/privateSquare/go/nexus-repository-cli/model"
	u "com/privateSquare/go/nexus-repository-cli/utils"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// CheckRepoExist checks if a repository exists in Nexus
func CheckRepoExist(nexusURL, repoId string, user m.AuthUser, verbose bool) bool {
	url := fmt.Sprintf("%s/service/local/repositories/%s", nexusURL, repoId)

	req := u.CreateBaseRequest("GET", url, nil, user, verbose)
	_, status := u.HTTPRequest(user, req, verbose)

	if status == "404 Not Found" {
		return false
	} else {
		return true
	}
}

func CheckGroupRepoExist(nexusURL, repoId string, user m.AuthUser, verbose bool) bool {
	url := fmt.Sprintf("%s/service/local/repo_groups/%s", nexusURL, repoId)

	req := u.CreateBaseRequest("GET", url, nil, user, verbose)
	_, status := u.HTTPRequest(user, req, verbose)

	if status == "404 Not Found" {
		return false
	} else {
		return true
	}
}

// CheckRepoId validates that repoId is not null
func CheckRepoId(repoId string) {
	if repoId == "" {
		log.Fatal("repoId is a required parameter")
	}
}

func AddRepoToGroup(nexusURL, groupRepoId, repositories string, user m.AuthUser, verbose bool) {
	CheckRepoId(groupRepoId)
	if repositories == "" {
		log.Fatal("repositories is a required paramter for adding a repo to a group")
	}
	url := fmt.Sprintf("%s/service/local/repo_groups/%s", nexusURL, groupRepoId)

	req := u.CreateBaseRequest("GET", url, nil, user, verbose)
	resp, status := u.HTTPRequest(user, req, verbose)

	var repository m.GroupRepository
	json.Unmarshal(resp, &repository)

	fmt.Println(repository, status)

	repoIdArray := strings.Split(repositories, ",")

	for _, repoId := range repoIdArray {
		if CheckRepoExist(nexusURL, repoId, user, verbose) {
			repository.Data.Repositories = append(repository.Data.Repositories, m.Repositories{ID: repoId})
		} else {
			log.Printf("Repository with ID=%s does not exist in Nexus, hence not adding it to the group repository", repoId)
		}

	}

	body, err := json.Marshal(repository)
	u.Error(err, "Error creating the request body")

	req = u.CreateBaseRequest("PUT", url, body, user, verbose)
	_, _ = u.HTTPRequest(user, req, verbose)

}
