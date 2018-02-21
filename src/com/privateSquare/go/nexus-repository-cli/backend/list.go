package backend

import (
	m "com/privateSquare/go/nexus-repository-cli/model"
	u "com/privateSquare/go/nexus-repository-cli/utils"
	"encoding/json"
	"strings"
)

/* Lists repositories available in nexus based on the input parameters
   returns a array of repositories*/
func List(nexusUrl, repoType, provider, repoPolicy string, user m.AuthUser, verbose bool) []string {
	url := nexusUrl + "/service/local/all_repositories"
	CheckRepoType(repoType)
	CheckProvider(provider)
	CheckMavenRepoPolicy(repoPolicy)
	var repositories []string
	req := u.CreateBaseRequest("GET", url, nil, user, verbose)
	response, _ := u.HTTPRequest(user, req, verbose)
	var jsonObject m.Repository
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
