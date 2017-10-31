package backend

import (
	"com/privateSquare/nexus-repository-cli/model"
	"com/privateSquare/nexus-repository-cli/utils"
	"encoding/json"
	"strings"
)

func List(nexusUrl, repoType, provider, repoPolicy string, user model.User, verbose bool) []string {
	url := nexusUrl + "/service/local/all_repositories"

	CheckRepoType(repoType)
	CheckProvider(provider)
	CheckMavenRepoPolicy(repoPolicy)

	var repositories []string
	response, _ := utils.HttpRequest(url, "GET", nil, user.Username, user.Password, verbose)

	var jsonObject model.Repository
	json.Unmarshal(response, &jsonObject)

	for _, repo := range jsonObject.Data {
		if repoType == "" && provider == "" && repoPolicy == "" {
			repositories = append(repositories, repo.Name)
		} else if repo.RepoType == repoType && provider == "" && repoPolicy == "" {
			repositories = append(repositories, repo.Name)
		} else if repoType == "" && repo.Format == provider && repoPolicy == "" {
			repositories = append(repositories, repo.Name)
		} else if repo.RepoType == repoType && repo.Format == provider && repoPolicy == "" {
			repositories = append(repositories, repo.Name)
		} else if repo.RepoType == repoType && repo.Format == provider && strings.ToUpper(repoPolicy) == repo.RepoPolicy {
			repositories = append(repositories, repo.Name)
		}
	}
	return repositories
}
