package backend

import (
	m "com/privateSquare/go/nexus-repository-cli/model"
	"fmt"
	"strings"
	"log"
	"os"
)

func JavaOnboarding(nexusURL, appSystemLetters string, user m.AuthUser, verbose bool){
	snapshotsRepo := "java-snapshots"
	releasesRepo := "java-releases"
	sharedRepo := "java-shared"
	thirdPartyRepo := "java-third-party"
	groupRepo := "java-group"
	if repoExists(nexusURL, releasesRepo, user, verbose) && repoExists(nexusURL, snapshotsRepo, user, verbose){
		//create repository target
		repoTargetName := fmt.Sprintf("%s-java", strings.ToLower(appSystemLetters))
		patternExpressions := fmt.Sprintf("/|/com/|/com/abnamro/|/com/abnamro/%s/.*", strings.ToLower(appSystemLetters))
		CreateMavenRepoTarget(nexusURL, repoTargetName, patternExpressions, user, verbose)
		//create privileges
		privilegeNameSnapshots := fmt.Sprintf("%s-maven-snapshots", strings.ToLower(appSystemLetters))
		privilegeNameReleases := fmt.Sprintf("%s-maven-releases", strings.ToLower(appSystemLetters))
		CreatePrivileges(nexusURL, privilegeNameSnapshots, snapshotsRepo, repoTargetName, user, verbose)
		CreatePrivileges(nexusURL, privilegeNameReleases, releasesRepo, repoTargetName, user, verbose)
		//create roles
		readerRole := fmt.Sprintf("%s_NEXUS_READERS", strings.ToUpper(appSystemLetters))
		publisherRole := fmt.Sprintf("%s_NEXUS_PUBLISHERS", strings.ToUpper(appSystemLetters))
		viewExt := "(view)"
		createExt := "(create)"
		readExt := "(read)"
		updateExt := "(update)"
		readPrivileges := fmt.Sprintf("%s - %s,%s - %s,%s - %s,%s - %s,%s - %s,%s - %s,%s - %s,%s - %s,%s - %s,%s - %s", snapshotsRepo, viewExt, releasesRepo, viewExt, sharedRepo, viewExt, thirdPartyRepo, viewExt, groupRepo, viewExt, sharedRepo, readExt, thirdPartyRepo, readExt, groupRepo, readExt, privilegeNameSnapshots, readExt, privilegeNameReleases, readExt)
		readRoles := ""
		publishPrivileges := fmt.Sprintf("%s - %s,%s - %s,%s - %s,%s - %s", privilegeNameSnapshots, createExt, privilegeNameSnapshots, updateExt, privilegeNameReleases, createExt, privilegeNameReleases, updateExt)
		publishRoles := fmt.Sprintf("%s", readerRole)
		CreateRole(nexusURL, readerRole, readPrivileges, readRoles, user, verbose)
		CreateRole(nexusURL, publisherRole, publishPrivileges, publishRoles, user, verbose)
	}else {
		log.Printf("'%s' and '%s' repositories are required for onboarding a java application to Nexus.", snapshotsRepo, releasesRepo)
		os.Exit(1)
	}
}

func JavaOffboarding(nexusURL, appSystemLetters string, user m.AuthUser, verbose bool){
	repoTargetName := fmt.Sprintf("%s-java", strings.ToLower(appSystemLetters))
	privilegeNameSnapshots := fmt.Sprintf("%s-maven-snapshots", strings.ToLower(appSystemLetters))
	privilegeNameReleases := fmt.Sprintf("%s-maven-releases", strings.ToLower(appSystemLetters))
	readerRole := fmt.Sprintf("%s_NEXUS_READERS", strings.ToUpper(appSystemLetters))
	publisherRole := fmt.Sprintf("%s_NEXUS_PUBLISHERS", strings.ToUpper(appSystemLetters))

	DeleteRole(nexusURL, readerRole, user, verbose)
	DeleteRole(nexusURL, publisherRole, user, verbose)
	DeletePrivileges(nexusURL, privilegeNameSnapshots, user, verbose)
	DeletePrivileges(nexusURL, privilegeNameReleases, user, verbose)
	DeleteRepoTarget(nexusURL, repoTargetName, user, verbose)
}