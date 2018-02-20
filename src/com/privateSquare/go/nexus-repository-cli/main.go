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
	///repository
	list := flag.Bool("list", false, "List the repositories in Nexus. Optional parameters: repoType, repoPolicy")
	createMavenHostedRepo := flag.Bool("createMavenHostedRepo", false, "Create a maven hosted repository (By default a snapshot repository is created). Required parameters: repoId Optional parameter: release (creates a release repository).")
	createMavenProxyRepo := flag.Bool("createMavenProxyRepo", false, "Create a maven proxy repository. Required parameters: repoId, remoteStorageURL. Optional parameters: exposed, browseable.")
	createMavenGroupRepo := flag.Bool("createMavenGroupRepo", false, "Create a maven group repository. Required parameters: repoId.")
	deleteRepo := flag.Bool("delete", false, "Delete a repository in Nexus. Required parameter: repoId.")
	addRepoToGroup := flag.Bool("addRepoToGroup", false, "Add a reposirory to a group repository. Required parameters: repoId, repositories.")
	///repository targets
	createMavenTarget := flag.Bool("createMavenTarget", false, "Create a maven repository target. Required parameters: repoTargetName, patternExpression.")
	deleteTarget := flag.Bool("deleteTarget", false, "Delete a repository target. Required parameters: repoTargetName.")
	///privileges
	createPrivileges := flag.Bool("createPrivileges", false, "Create repository privileges. Required parameters: privilegeName, repoTargetName.")
	deletePrivileges := flag.Bool("deletePrivileges", false, "Delete repository privileges. Required parameters: privilegeName.")
	///roles
	createRole := flag.Bool("createRole", false, "Create roles. Required parameters: roleName, privileges, roles.")
	deleteRole := flag.Bool("deleteRole", false, "Delete roles. Required parameters: roleName.")
	roleName := flag.String("roleName", "", "Role name.")
	privileges := flag.String("privileges","","Comma separated privilege name values.")
	roles := flag.String("roles","","Comma separated role name values.")
	//variables
	///general
	username := flag.String("username", "", "Username for authentication.")
	password := flag.String("password", "", 	"Password for authentication.")
	nexusURL := flag.String("nexusURL", "http://localhost:8081/nexus", "Nexus server URL.")
	verbose := flag.Bool("verbose", false, "Set this flag for Debug logs.")
	///repository
	repoId := flag.String("repoId", "", "ID of a Repository.")
	repoType := flag.String("repoType", "", "Type of a repository. Possible values : hosted/proxy/group.")
	repoPolicy := flag.String("repoPolicy", "", "Policy of the hosted repository. Possible values : snapshot/release.")
	provider := flag.String("provider", "", "Repository provider. Possible values: maven2/npm/nuget.")
	exposed := flag.Bool("exposed", false, "Set this flag to expose the repository in nexus.")
	browseable := flag.Bool("browseable", false, "Set this flag to make the repository browseable in nexus.")
	release := flag.Bool("release", false, "Set this flag for creating a maven release repository.")
	repositories := flag.String("repositories", "", "Comma separated value of repositories to be added to a group.")
	remoteStorageURL := flag.String("remoteStorageURL", "", "Remote storage url to proxy in Nexus.")
	///repository targets
	repoTargetName := flag.String("repoTargetName", "", "Repository target name.")
	pattern := flag.String("pattern", "", "Repository target pattern expression. Can be comma separated values.")
	///privileges
	privilegeName := flag.String("privilegeName", "", "Repository Privilege name.")
	///roles

	flag.Parse()

	user := m.AuthUser{*username, *password}

	if *username == "" || *password == "" {
		log.Fatal("username and password are required parameters")
	} else if *nexusURL == "" {
		log.Fatal("nexusUrl is a required parameter")
	}

	//TODO : Should developers have delete access on their released artifacts?

	if *list == true {
		repositories := b.List(*nexusURL, *repoType, *provider, *repoPolicy, user, *verbose)
		u.PrintStringArray(repositories)
		fmt.Printf("No of %s %s repositories in Nexus : %d", *provider, *repoType, len(repositories))
	} else if *createMavenHostedRepo == true {
		b.CreateMavenHostedRepo(*nexusURL, *repoId, user, *release, *verbose)
	} else if *createMavenProxyRepo == true {
		b.CreateMavenProxyRepo(*nexusURL, *repoId, *remoteStorageURL, user, *exposed, *browseable, *verbose)
	} else if *createMavenGroupRepo == true {
		b.CreateMavenGroupRepo(*nexusURL, *repoId, user, *verbose)
	} else if *addRepoToGroup == true {
		b.AddRepoToGroup(*nexusURL, *repoId, *repositories, user, *verbose)
	}else if *createMavenTarget == true {
		b.CreateMavenRepoTarget(*nexusURL, *repoTargetName, *pattern, user, *verbose)
	} else if *deleteTarget == true {
		b.DeleteRepoTarget(*nexusURL, *repoTargetName,  user, *verbose)
	} else if *createPrivileges == true {
		b.CreatePrivileges(*nexusURL, *privilegeName, *repoTargetName, user, *verbose)
	} else if *deletePrivileges == true {
		b.DeletePrivileges(*nexusURL, *privilegeName, user, *verbose)
	} else if *createRole == true {
		b.CreateRole(*nexusURL, *roleName, *privileges, *roles, user, *verbose)
	} else if *deleteRole == true {
		b.DeleteRole(*nexusURL, *roleName, user, *verbose)
	} else if *deleteRepo == true {
		b.CheckRepoId(*repoId)
		b.DeleteRepo(user, *nexusURL, *repoId, *verbose)
	} else {
		flag.Usage()
		log.Fatal("Select a valid action flag")
	}
}
