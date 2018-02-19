package backend

import (
	"fmt"
	u "com/privateSquare/go/nexus-repository-cli/utils"
	m "com/privateSquare/go/nexus-repository-cli/model"
	"encoding/json"
	"log"
	"os"
	"strings"
	"regexp"
)

func getRepoTargetID (nexusURL, repoTargetName string, user m.AuthUser, verbose bool) string{
	url := fmt.Sprintf("%s/service/local/repo_targets", nexusURL)
	req := u.CreateBaseRequest("GET", url, nil, user, verbose)
	respBody, _ := u.HTTPRequest(user, req, verbose)
	var (
		repoTargetId string
		repoTargets m.RepositoryTargets
	)
	json.Unmarshal(respBody, &repoTargets)
	for _, repoTarget := range repoTargets.Data{
		if repoTarget.Name == repoTargetName {
			repoTargetId = repoTarget.ID
		}
	}
	return repoTargetId
}

func repoTargetExists (nexusURL, repoTargetName string, user m.AuthUser, verbose bool) bool{
	var isExists bool
	repoTargetID := getRepoTargetID(nexusURL, repoTargetName, user, verbose)
	if repoTargetID != ""{
		isExists = true
	} else {
		isExists = false
	}
	return isExists
}

func CreateMavenRepoTarget(nexusURL, repoTargetName, patternExpressions string, user m.AuthUser, verbose bool){
	if repoTargetName == "" || patternExpressions == "" {
		log.Println("repoTargetName and patternExpressions are required parameters for creating a repository target.")
		os.Exit(1)
	}
	createRepoTarget(nexusURL, repoTargetName, patternExpressions, "maven2", user, verbose)
}

func createRepoTarget(nexusURL, repoTargetName, patternExpressions, contentClass string ,user m.AuthUser, verbose bool){
	if !repoTargetExists(nexusURL, repoTargetName, user, verbose){
		url := fmt.Sprintf("%s/service/local/repo_targets", nexusURL)
		patternExpressionsArray := strings.Split(patternExpressions, ",")
		repoTarget := m.RepositoryTargetCreate{
			Data: m.RepositoryTargetCreateData{
				Name:repoTargetName,
				ContentClass:contentClass,
				Patterns: patternExpressionsArray,
			},
		}
		body, err := json.Marshal(repoTarget)
		u.Error(err, "Error creating the request body")
		req := u.CreateBaseRequest("POST", url, body, user, verbose)
		_, status := u.HTTPRequest(user, req, verbose)
		if status == "201 Created"{
			log.Printf("Repository target '%s' is created.", repoTargetName)
		}else{
			log.Printf("Repository target create request failed with status : %s.", status)
		}
	} else {
		log.Printf("Repository target '%s' already exists.", repoTargetName)
		os.Exit(1)
	}
}

func DeleteRepoTarget(nexusURL, repoTargetName string ,user m.AuthUser, verbose bool){
	if repoTargetName == "" {
		log.Println("repoTargetName is a required paramteter for deleting a repository target.")
		os.Exit(1)
	}
	if repoTargetExists(nexusURL, repoTargetName, user, verbose){
		repoTargetID := getRepoTargetID(nexusURL, repoTargetName, user, verbose)
		url := fmt.Sprintf("%s/service/local/repo_targets/%s", nexusURL, repoTargetID)
		req := u.CreateBaseRequest("DELETE", url, nil, user, verbose)
		u.HTTPRequest(user, req, verbose)
		log.Printf("Repository target for '%s' deleted.\n", repoTargetName)
	}else {
		log.Printf("Repository target '%s' does not exists.", repoTargetName)
		os.Exit(1)
	}
}

func getRepoPrivilegesID (nexusURL, repoPrivilegeName string, user m.AuthUser, verbose bool) []string{
	url := fmt.Sprintf("%s/service/local/privileges", nexusURL)
	req := u.CreateBaseRequest("GET", url, nil, user, verbose)
	respBody, _ := u.HTTPRequest(user, req, verbose)
	var (
		repoPrivileges m.RepoPrivileges
		repoPrivilegesID []string
	)
	json.Unmarshal(respBody, &repoPrivileges)
	cRegexp, err := regexp.Compile(repoPrivilegeName + " ")
	u.Error(err, "There was a error compiling the regex.")
	for _, repoPrivilege := range repoPrivileges.Data{
		if cRegexp.MatchString(repoPrivilege.Name) {
			repoPrivilegesID = append(repoPrivilegesID, repoPrivilege.ID)
		}
	}
	return repoPrivilegesID
}

func repoPrivilegesExists (nexusURL, repoPrivilegeName string, user m.AuthUser, verbose bool) bool{
	var isExists bool
	repoPrivileges := getRepoPrivilegesID(nexusURL, repoPrivilegeName, user, verbose)
	if len(repoPrivileges) > 0 {
		isExists = true
	}else {
		isExists = false
	}
	return isExists
}

func CreateRepoPrivileges(nexusURL, repoPrivilegeName, repoTargetName string ,user m.AuthUser, verbose bool){
	if repoPrivilegeName == "" || repoTargetName == "" {
		log.Println("repoPrivilegeName and repoTargetName are required parameters for creating repository privileges.")
		os.Exit(1)
	}
	if !repoPrivilegesExists(nexusURL, repoPrivilegeName, user, verbose){
		url := fmt.Sprintf("%s/service/local/privileges_target", nexusURL)
		repoPrivilege := m.RepoPrivilegesCreate{
			Data: m.RepoPrivilegeCreateData{
				Name:repoPrivilegeName,
				Description:repoPrivilegeName,
				RepositoryID:"",
				RepositoryGroupID:"",
				Type:"target",
				RepositoryTargetID:getRepoTargetID(nexusURL, repoTargetName, user, verbose),
				Method: []string{"create", "read", "update", "delete"},
			},
		}
		body, err := json.Marshal(repoPrivilege)
		u.Error(err, "Error creating the request body")
		req := u.CreateBaseRequest("POST", url, body, user, verbose)
		u.HTTPRequest(user, req, verbose)
		log.Printf("CRUD Repository privileges for '%s' created.\n", repoPrivilegeName)
	}else {
		log.Printf("Repository privileges for '%s' already exists.", repoPrivilegeName)
		os.Exit(1)
	}
}

func DeleteRepoPrivileges(nexusURL, repoPrivilegeName string ,user m.AuthUser, verbose bool){
	if repoPrivilegeName == "" {
		log.Println("repoPrivilegeName is a required parameter for deleting repository privileges.")
		os.Exit(1)
	}
	if repoPrivilegesExists(nexusURL, repoPrivilegeName, user, verbose){
		repoPrivilegesID := getRepoPrivilegesID(nexusURL, repoPrivilegeName, user, verbose)
		for _, repoPrivilegeID := range repoPrivilegesID{
			url := fmt.Sprintf("%s/service/local/privileges/%s", nexusURL, repoPrivilegeID)
			req := u.CreateBaseRequest("DELETE", url, nil, user, verbose)
			u.HTTPRequest(user, req, verbose)
		}
		log.Printf("All Repository privileges for '%s' deleted.\n", repoPrivilegeName)
	}else{
		log.Printf("Repository privileges for '%s' does not exists.", repoPrivilegeName)
		os.Exit(1)
	}
}
