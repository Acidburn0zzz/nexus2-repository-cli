package backend

import (
	m "com/privateSquare/go/nexus-repository-cli/model"
	u "com/privateSquare/go/nexus-repository-cli/utils"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// repoExists checks if a repository exists in Nexus
func repoExists(nexusURL, repoID string, user m.AuthUser, verbose bool) bool {
	url := fmt.Sprintf("%s/service/local/repositories/%s", nexusURL, repoID)
	req := u.CreateBaseRequest("GET", url, nil, user, verbose)
	_, status := u.HTTPRequest(user, req, verbose)
	if status == "200 OK" {
		return true
	} else {
		return false
	}
}

func groupRepoExists(nexusURL, repoID string, user m.AuthUser, verbose bool) bool {
	url := fmt.Sprintf("%s/service/local/repo_groups/%s", nexusURL, repoID)
	req := u.CreateBaseRequest("GET", url, nil, user, verbose)
	_, status := u.HTTPRequest(user, req, verbose)
	if status == "200 OK" {
		return true
	} else {
		return false
	}
}

// CheckRepoId validates that repoID is not null
func checkRepoId(repoID string) {
	if repoID == "" {
		log.Fatal("repoID is a required parameter")
	}
}

func AddRepoToGroup(nexusURL, groupRepoId, repositories string, user m.AuthUser, verbose bool) {
	checkRepoId(groupRepoId)
	if repositories == "" {
		log.Fatal("repositories is a required paramter for adding a repo to a group")
	}
	url := fmt.Sprintf("%s/service/local/repo_groups/%s", nexusURL, groupRepoId)
	req := u.CreateBaseRequest("GET", url, nil, user, verbose)
	resp, _ := u.HTTPRequest(user, req, verbose)
	var repository m.GroupRepository
	json.Unmarshal(resp, &repository)
	repoIDArray := strings.Split(repositories, ",")
	for _, repoID := range repoIDArray {
		if repoExists(nexusURL, repoID, user, verbose) {
			repository.Data.Repositories = append(repository.Data.Repositories, m.Repositories{ID: repoID})
			log.Printf("Adding repository with ID=%s to the group repository '%s'.\n", repoID, groupRepoId)
		} else {
			log.Printf("Repository with ID=%s does not exist in Nexus, hence not adding it to the group repository.\n", repoID)
		}

	}
	body, err := json.Marshal(repository)
	u.Error(err, "Error creating the request body")
	req = u.CreateBaseRequest("PUT", url, body, user, verbose)
	u.HTTPRequest(user, req, verbose)
}
