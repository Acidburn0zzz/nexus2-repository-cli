package main

import (
	"com/privateSquare/nexus-repository-cli/model"
	"com/privateSquare/nexus-repository-cli/backend"
	"com/privateSquare/nexus-repository-cli/utils"
	"flag"
	"log"
	"fmt"
)

func main() {
	//actions
	list := flag.Bool("list", false, "List the repositories in Nexus. Optional parameters: repoType, repoPolicy")
	create := flag.Bool("create", false, "Create a repository in Nexus. Required parameter: repoId, repoType, provider, repoPolicy (only for maven2). Optional parameter: exposed")
	delete := flag.Bool("delete", false, "Delete a repository in Nexus. Required parameter: repoId")
	addRepoToGroup := flag.Bool("addRepoToGroup", false, "Add a reposirory to a group repository. Required paramters: repoId, repositories ")
	//variables
	username := flag.String("username", "", "Username for authentication")
	password := flag.String("password", "", "Password for authentication")
	nexusUrl := flag.String("nexusUrl", "http://localhost:8081/nexus", "Nexus server URL")
	repoId := flag.String("repoId", "", "ID of the Repository")
	repoType := flag.String("repoType", "", "Type of a repository. Possible values : hosted/proxy/group")
	repoPolicy := flag.String("repoPolicy", "", "Policy of the hosted repository. Possible values : snapshot/release")
	provider := flag.String("provider", "", "Repository provider. Possible values: maven2/npm/nuget")
	repositories := flag.String("repositories", "", "Comma separated value of repositories to be added to a group.")
	remoteStorageUrl := flag.String("remoteStorageUrl", "", "Remote storage url to proxy in Nexus")
	exposed := flag.Bool("exposed", false, "Set this flag to expose the repository in nexus.")
	verbose := flag.Bool("verbose", false, "Set this flag for Debug logs.")
	flag.Parse()

	user := model.User{*username, *password}

	if *username == "" || *password == "" {
		log.Fatal("username and password are required parameters")
	} else if *nexusUrl == "" {
		log.Fatal("nexusUrl is a required parameter")
	}

	utils.PrintHelp()

	if *list == true {
		repositories := backend.List(*nexusUrl, *repoType,*provider, *repoPolicy,  user, *verbose)
		utils.PrintStringArray(repositories)
		fmt.Printf("No of %s %s repositories in Nexus : %d", *provider, *repoType, len(repositories))
	}else if *create == true {
		if *repoId != "" && *repoType != "" && *provider != "" {
			backend.CheckRepoType(*repoType)
			backend.CheckProvider(*provider)
			switch *repoType {
			case "hosted":
				if *provider == "maven2" {
					if *repoPolicy == ""{
						log.Fatal("repoPolicy is a required parameter for creating a hosted maven repository in Nexus")
					}
					backend.CheckMavenRepoPolicy(*repoPolicy)
					backend.CreateHostedRepo(user, *nexusUrl, *repoId, *repoType, *repoPolicy, "maven2", *exposed, *verbose)
				}
				if *provider == "npm" {
					backend.CreateHostedRepo(user, *nexusUrl, *repoId, *repoType, "mixed", "npm-hosted", *exposed, *verbose)
				}
				if *provider == "nuget" {
					backend.CreateHostedRepo(user, *nexusUrl, *repoId, *repoType, "mixed", "nuget-proxy", *exposed, *verbose)
				}
			case "proxy":
				if *provider == "maven2" && *remoteStorageUrl != "" {
					backend.CreateProxyRepo( user, *nexusUrl, *repoId, *repoType, "release", *remoteStorageUrl, "maven2", *exposed, *verbose)
				} else if *provider == "npm" && *remoteStorageUrl != "" {
					backend.CreateProxyRepo( user, *nexusUrl, *repoId, *repoType, "mixed", *remoteStorageUrl, "npm-proxy", *exposed, *verbose)
				} else if *provider == "nuget" && *remoteStorageUrl != "" {
					backend.CreateProxyRepo( user, *nexusUrl, *repoId, *repoType, "mixed", *remoteStorageUrl, "nuget-proxy", *exposed, *verbose)
				} else {
					log.Fatal("remoteStorageUrl is a required parameter for creating a proxy repository")
				}
			case "group":
				if *provider == "maven2" && *repositories != "" {
					backend.CreateGroupRepo(user, *nexusUrl, *repoId, *repoType, *repositories, "maven2", *exposed, *verbose)
				} else if *provider == "npm" && *repositories != "" {
					backend.CreateGroupRepo(user, *nexusUrl, *repoId, *repoType, *repositories, "npm-group", *exposed, *verbose)
				} else if *provider == "nuget" && *repositories != "" {
					backend.CreateGroupRepo(user, *nexusUrl, *repoId, *repoType, *repositories, "nuget-group", *exposed, *verbose)
				} else {
					log.Fatal("repositories is a required parameter for creating a group repository")
				}
			}
		}else {
			log.Fatal("repoId ,repoType and provider are required parameters for creating a repository in Nexus")
		}
	} else if *delete == true {
		if *repoId == "" {
			log.Fatal("repoId is a required parameter")
		}
		backend.DeleteRepo(user, *nexusUrl, *repoId, *verbose)
	} else if *addRepoToGroup == true {
		fmt.Println("Not implemented yet")
	} else {
		flag.Usage()
		log.Fatal("Select a valid action flag")
	}
}