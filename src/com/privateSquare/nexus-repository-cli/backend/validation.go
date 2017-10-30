package backend

import (
	"log"
	"com/privateSquare/nexus-repository-cli/utils"
	"com/privateSquare/nexus-repository-cli/model"
)

func CheckRepoExist(user model.User, nexusUrl, repoId string, verbose bool) bool{

	url := nexusUrl + "/service/local/repositories/" + repoId

	_, status := utils.HttpRequest(url, "GET", nil, user.Username, user.Password, verbose)

	if status == "404 Not Found" {
		return false
	}else {
		return true
	}
}

func CheckRepoType (repoType string){
	if repoType != "" && repoType != "hosted" && repoType != "group" && repoType != "proxy" {
		log.Fatal("repoType value is invalid. Possible values are hosted/group/proxy")
	}
}

func CheckProvider (provider string){
	if provider != "" && provider != "maven2" && provider != "npm" && provider != "nuget" {
		log.Fatal("provider value is invalid. Possible values are maven2/npm/nuget")
	}
}

func CheckMavenRepoPolicy (repoPolicy string){
	if repoPolicy != "" && repoPolicy != "release" && repoPolicy != "snapshot" {
		log.Fatal("repoPolicy value is invalid. Possible values are snapshot/release")
	}
}