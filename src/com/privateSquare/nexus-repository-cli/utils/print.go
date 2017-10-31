package utils

import (
	"flag"
	"fmt"
	"os"
)

func PrintStringArray(stringArray []string) {
	for _, array := range stringArray {
		fmt.Println(array)
	}

}

func PrintHelp() {

	helpString := `
 Usage: ./nexus-repository-cli.exe [option] [parameters...]

 [options]
  -list
        List the repositories in Nexus. Optional parameters: repoType, repoPolicy
  -create
        Create a repository in Nexus. Required parameter: repoId, repoType, provider, repoPolicy (only for maven2). Optional parameter: exposed
  -delete
        Delete a repository in Nexus. Required parameter: repoId
  -addRepoToGroup
        Add a reposirory to a group repository. Required paramters: repoId, repositories

 [parameters]
  -nexusUrl string
        Nexus server URL (default "http://localhost:8081/nexus")
  -exposed
        Set this flag to expose the repository in nexus.
  -username string
        Username for authentication
  -password string
        Password for authentication
  -repoId string
        ID of the Repository
  -repoType string
        Type of a repository. Possible values : hosted/proxy/group
  -repoPolicy string
        Policy of the hosted repository. Possible values : snapshot/release
  -provider string
        Repository provider. Possible values: maven2/npm/nuget
  -remoteStorageUrl string
        Remote storage url to proxy in Nexus
  -repositories string
        Comma separated value of repositories to be added to a group.
  -verbose
        Set this flag for Debug logs.
	`

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, helpString)
	}
}
