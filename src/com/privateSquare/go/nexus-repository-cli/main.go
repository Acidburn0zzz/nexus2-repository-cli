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
	///onboarding
	javaOnboarding := flag.Bool("javaOnboarding", false, "Create a new space for a application in the java repo structure. Required paramters: appSysLetters.")
	///repository
	list := flag.Bool("list", false, "List the repositories in Nexus. Optional parameters: repoType, repoPolicy")
	createMavenHostedRepo := flag.Bool("createMavenHostedRepo", false, "Create a maven hosted repository (By default a snapshot repository is created). Required parameters: repoID Optional parameter: release (creates a release repository).")
	createMavenProxyRepo := flag.Bool("createMavenProxyRepo", false, "Create a maven proxy repository. Required parameters: repoID, remoteStorageURL. Optional parameters: exposed, browseable.")
	createMavenGroupRepo := flag.Bool("createMavenGroupRepo", false, "Create a maven group repository. Required parameters: repoID.")
	deleteRepo := flag.Bool("deleteRepo", false, "Deletes a hosted/proxy repository. Required parameter: repoID.")
	deleteGroupRepo := flag.Bool("deleteGroupRepo", false, "Deletes a group repository. Required parameter: repoID.")
	addRepoToGroup := flag.Bool("addRepoToGroup", false, "Add a reposirory to a group repository. Required parameters: repoID, repositories.")
	///repository targets
	createMavenTarget := flag.Bool("createMavenTarget", false, "Create a maven repository target. Required parameters: repoTargetName, patternExpression.")
	deleteTarget := flag.Bool("deleteTarget", false, "Delete a repository target. Required parameters: repoTargetName.")
	///privileges
	createPrivileges := flag.Bool("createPrivileges", false, "Create repository privileges. Required parameters: privilegeName, repoTargetName.")
	deletePrivileges := flag.Bool("deletePrivileges", false, "Delete repository privileges. Required parameters: privilegeName.")
	///roles
	createRole := flag.Bool("createRole", false, "Create roles. Required parameters: roleName, privileges, roles.")
	deleteRole := flag.Bool("deleteRole", false, "Delete roles. Required parameters: roleName.")
	//variables
	///general
	username := flag.String("username", "", "Username for authentication.")
	password := flag.String("password", "", 	"Password for authentication.")
	nexusURL := flag.String("nexusURL", "http://localhost:8081/nexus", "Nexus server URL.")
	appSysLetters := flag.String("appSysLetters", "", "Applictaion system letters of an application. (3 or 4 letter code)")
	verbose := flag.Bool("verbose", false, "Set this flag for Debug logs.")
	///repository
	repoID := flag.String("repoID", "", "ID of a Repository.")
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
	roleName := flag.String("roleName", "", "Role name.")
	privileges := flag.String("privileges","","Comma separated privilege name values.")
	roles := flag.String("roles","","Comma separated role name values.")
	flag.Parse()

	user := m.AuthUser{*username, *password}

	if *username == "" || *password == "" {
		log.Fatal("username and password are required parameters")
	} else if *nexusURL == "" {
		log.Fatal("nexusUrl is a required parameter")
	}

	//b.JavaOnboarding(*nexusURL, "ABC", user, *verbose)

	if * javaOnboarding == true {
		b.JavaOnboarding(*nexusURL, *appSysLetters, user, *verbose)
	} else if *list == true {
		repositories := b.List(*nexusURL, *repoType, *provider, *repoPolicy, user, *verbose)
		u.PrintStringArray(repositories)
		fmt.Printf("No of %s %s repositories in Nexus : %d", *provider, *repoType, len(repositories))
	} else if *createMavenHostedRepo == true {
		b.CreateMavenHostedRepo(*nexusURL, *repoID, user, *release, *verbose)
	} else if *createMavenProxyRepo == true {
		b.CreateMavenProxyRepo(*nexusURL, *repoID, *remoteStorageURL, user, *exposed, *browseable, *verbose)
	} else if *createMavenGroupRepo == true {
		b.CreateMavenGroupRepo(*nexusURL, *repoID, *repositories, user, *verbose)
	} else if *addRepoToGroup == true {
		b.AddRepoToGroup(*nexusURL, *repoID, *repositories, user, *verbose)
	} else if *deleteRepo == true {
		b.DeleteRepo(*nexusURL, *repoID, user, *verbose)
	} else if *deleteGroupRepo == true {
		b.DeleteGroupRepo(*nexusURL, *repoID, user, *verbose)
	} else if *createMavenTarget == true {
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
	} else {
		//flag.Usage()
		log.Fatal("Select a valid action flag")
	}
}
