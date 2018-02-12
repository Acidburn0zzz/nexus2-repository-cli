package backend

import (
	m "com/privateSquare/go/nexus-repository-cli/model"
	u "com/privateSquare/go/nexus-repository-cli/utils"
	"encoding/json"
)

// DeleteRepo deletes a repository in Nexus
func DeleteRepo(user m.AuthUser, nexusUrl, repoId string, verbose bool) {

	url := nexusUrl + "/service/local/all_repositories"

	req := u.CreateBaseRequest("GET", url, nil, user, verbose)
	response, _ := u.HTTPRequest(user, req, verbose)

	var repoType string
	var jsonObject m.Repository
	json.Unmarshal(response, &jsonObject)

	for _, repo := range jsonObject.Data {
		if repo.ID == repoId {
			repoType = repo.RepoType
		}
	}
	if repoType == "hosted" || repoType == "proxy" {
		//url := nexusUrl + "/service/local/repositories/" + repoId
		//_, status := utils.HttpRequest(url, "DELETE", nil, user.Username, user.Password, verbose)
		//log.Println(status)
	} else if repoType == "group" {
		//url := nexusUrl + "/service/local/repo_groups/" + repoId
		//_, status := utils.HttpRequest(url, "DELETE", nil, user.Username, user.Password, verbose)
		//log.Println(status)
	} else {
		//log.Printf("Repository with ID=%s does not exist.", repoId)
	}
}
