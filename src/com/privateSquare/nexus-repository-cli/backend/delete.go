package backend

import (
	"com/privateSquare/nexus-repository-cli/model"
	"com/privateSquare/nexus-repository-cli/utils"
	"encoding/json"
	"log"
)

// DeleteRepo deletes a repository in Nexus
func DeleteRepo(user model.User, nexusUrl, repoId string, verbose bool) {

	url := nexusUrl + "/service/local/all_repositories"
	response, _ := utils.HttpRequest(url, "GET", nil, user.Username, user.Password, verbose)

	var repoType string
	var jsonObject model.Repository
	json.Unmarshal(response, &jsonObject)

	for _, repo := range jsonObject.Data {
		if repo.ID == repoId {
			repoType = repo.RepoType
		}
	}
	if repoType == "hosted" || repoType == "proxy" {
		url := nexusUrl + "/service/local/repositories/" + repoId
		_, status := utils.HttpRequest(url, "DELETE", nil, user.Username, user.Password, verbose)
		log.Println(status)
	} else if repoType == "group" {
		url := nexusUrl + "/service/local/repo_groups/" + repoId
		_, status := utils.HttpRequest(url, "DELETE", nil, user.Username, user.Password, verbose)
		log.Println(status)
	} else {
		log.Printf("Repository with ID=%s does not exist.", repoId)
	}
}
