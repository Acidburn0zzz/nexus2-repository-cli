package backend

import (
	m "com/privateSquare/go/nexus-repository-cli/model"
	u "com/privateSquare/go/nexus-repository-cli/utils"
	"fmt"
	"log"
	"os"
)

// DeleteRepo deletes a hosted/proxy repository in Nexus
func DeleteRepo(nexusURL, repoID string, user m.AuthUser, verbose bool) {
	checkRepoId(repoID)
	if repoExists(nexusURL, repoID, user, verbose){
		url := fmt.Sprintf("%s/service/local/repositories/%s", nexusURL, repoID)
		req := u.CreateBaseRequest("DELETE", url, nil, user, verbose)
		u.HTTPRequest(user, req, verbose)
		log.Printf("Repository '%s' deleted.\n", repoID)
	} else{
		log.Printf("Repository '%s' does not exist.\n", repoID)
		os.Exit(1)
	}
}

func DeleteGroupRepo(nexusURL, repoID string, user m.AuthUser, verbose bool) {
	checkRepoId(repoID)
	if groupRepoExists(nexusURL, repoID, user, verbose){
		url := fmt.Sprintf("%s/service/local/repo_groups/%s", nexusURL, repoID)
		req := u.CreateBaseRequest("DELETE", url, nil, user, verbose)
		u.HTTPRequest(user, req, verbose)
		log.Printf("Repository group '%s' deleted.\n", repoID)
	} else{
		log.Printf("Repository group '%s' does not exist.\n", repoID)
		os.Exit(1)
	}
}