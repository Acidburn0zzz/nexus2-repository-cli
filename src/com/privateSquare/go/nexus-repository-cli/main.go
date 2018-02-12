package main

import (
	b "com/privateSquare/go/nexus-repository-cli/backend"
	m "com/privateSquare/go/nexus-repository-cli/model"
	u "com/privateSquare/go/nexus-repository-cli/utils"
	"flag"
	"fmt"
	"log"
)

func main() {
	//actions
	list := flag.Bool("list", false, "List the repositories in Nexus. Optional parameters: repoType, repoPolicy")
	//create := flag.Bool("create", false, "Create a repository in Nexus. Required parameter: repoId, repoType, provider, repoPolicy (only for maven2). Optional parameter: exposed")
	createMavenHostedRepo := flag.Bool("createMavenHostedRepo", false, "Create a maven hosted repository (By default a snapshot repository is created). Required parameters: repoId Optional parameter: release (creates a release repository)")
	createMavenProxyRepo := flag.Bool("createMavenProxyRepo", false, "Create a maven proxy repository. Required parameters: repoId, remoteStorageURL. Optional parameters: exposed, browseable")
	createMavenGroupRepo := flag.Bool("createMavenGroupRepo", false, "Create a maven group repository. Required parameters: repoId")

	delete := flag.Bool("delete", false, "Delete a repository in Nexus. Required parameter: repoId")
	addRepoToGroup := flag.Bool("addRepoToGroup", false, "Add a reposirory to a group repository. Required parameters: repoId, repositories ")

	//variables
	username := flag.String("username", "admin", "Username for authentication")
	password := flag.String("password", "admin123", "Password for authentication")
	nexusUrl := flag.String("nexusUrl", "http://localhost:8081/nexus", "Nexus server URL")
	repoId := flag.String("repoId", "", "ID of the Repository")
	repoType := flag.String("repoType", "", "Type of a repository. Possible values : hosted/proxy/group")
	repoPolicy := flag.String("repoPolicy", "", "Policy of the hosted repository. Possible values : snapshot/release")
	provider := flag.String("provider", "", "Repository provider. Possible values: maven2/npm/nuget")
	//repositories := flag.String("repositories", "", "Comma separated value of repositories to be added to a group.")
	remoteStorageURL := flag.String("remoteStorageURL", "", "Remote storage url to proxy in Nexus")
	exposed := flag.Bool("exposed", false, "Set this flag to expose the repository in nexus.")
	browseable := flag.Bool("browseable", false, "Set this flag to make the repository browseable in nexus.")
	release := flag.Bool("release", false, "Set this flag for creating a maven release repository.")
	verbose := flag.Bool("verbose", false, "Set this flag for Debug logs.")
	flag.Parse()

	user := m.AuthUser{*username, *password}

	if *username == "" || *password == "" {
		log.Fatal("username and password are required parameters")
	} else if *nexusUrl == "" {
		log.Fatal("nexusUrl is a required parameter")
	}

	//b.AddRepoToGroup(*nexusUrl, "ATS-g", *repositories, user, *verbose)

	if *list == true {
		repositories := b.List(*nexusUrl, *repoType, *provider, *repoPolicy, user, *verbose)
		u.PrintStringArray(repositories)
		fmt.Printf("No of %s %s repositories in Nexus : %d", *provider, *repoType, len(repositories))
	} else if *createMavenHostedRepo == true {
		b.CreateMavenHostedRepo(*nexusUrl, *repoId, user, *release, *verbose)
	} else if *createMavenProxyRepo == true {
		b.CreateMavenProxyRepo(*nexusUrl, *repoId, *remoteStorageURL, user, *exposed, *browseable, *verbose)
	} else if *createMavenGroupRepo == true {
		b.CreateMavenGroupRepo(*nexusUrl, *repoId, user, *verbose)
	} else if *delete == true {
		b.CheckRepoId(*repoId)
		b.DeleteRepo(user, *nexusUrl, *repoId, *verbose)
	} else if *addRepoToGroup == true {
		fmt.Println("Not implemented yet")
	} else {
		flag.Usage()
		log.Fatal("Select a valid action flag")
	}
}
