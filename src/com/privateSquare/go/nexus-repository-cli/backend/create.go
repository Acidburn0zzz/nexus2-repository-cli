package backend

import (
	m "com/privateSquare/go/nexus-repository-cli/model"
	u "com/privateSquare/go/nexus-repository-cli/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateMavenHostedRepo(nexusURL, repoID string, user m.AuthUser, release, verbose bool) {
	checkRepoId(repoID)
	var (
		repoPolicy  string
		writePolicy string
	)
	if release {
		repoPolicy = "RELEASE"
		writePolicy = "ALLOW_WRITE_ONCE"
	} else {
		repoPolicy = "SNAPSHOT"
		writePolicy = "ALLOW_WRITE"
	}
	createHostedRepo(nexusURL, repoID, "maven2", repoPolicy, writePolicy, user, verbose)
}

func CreateMavenProxyRepo(nexusURL, repoID, remoteStorageURL string, user m.AuthUser, exposed, browseable, verbose bool) {
	checkRepoId(repoID)
	if remoteStorageURL == "" {
		log.Fatal("remoteStorageURL is a required parameter for creating a proxy repository")
	}
	createProxyRepo(nexusURL, repoID, "maven2", "release", remoteStorageURL, user, exposed, browseable, verbose)
}

func CreateMavenGroupRepo(nexusURL, repoID, repositories string, user m.AuthUser, verbose bool) {
	checkRepoId(repoID)
	createGroupRepo(nexusURL, repoID, "group", "maven2", user, verbose)
	if repositories == "" {
		log.Println("repostories field is empty, hence creating a empty group repository.")
	}else{
		AddRepoToGroup(nexusURL, repoID, repositories, user, verbose)
	}
}

func CreateNPMHostedRepo(nexusURL, repoID string, user m.AuthUser, verbose bool) {
	checkRepoId(repoID)
	createHostedRepo(nexusURL, repoID, "npm-hosted", "mixed", "ALLOW_WRITE_ONCE", user, verbose)
}

func CreateNPMProxyRepo(nexusURL, repoID, remoteStorageURL string, user m.AuthUser, exposed, browseable, verbose bool) {
	checkRepoId(repoID)
	if remoteStorageURL == "" {
		log.Fatal("remoteStorageURL is a required parameter for creating a proxy repository")
	}
	createProxyRepo(nexusURL, repoID, "maven2", "release", remoteStorageURL, user, exposed, browseable, verbose)
}

func CreateNPMGroupRepo(nexusURL, repoID string, user m.AuthUser, verbose bool) {
	checkRepoId(repoID)
	createGroupRepo(nexusURL, repoID, "group", "npm-group", user, verbose)
}

func CreateNugetHostedRepo(nexusURL, repoID string, user m.AuthUser, verbose bool) {
	checkRepoId(repoID)
	createHostedRepo(nexusURL, repoID, "nuget-proxy", "mixed", "ALLOW_WRITE_ONCE", user, verbose)
}

func CreateNugetProxyRepo(nexusURL, repoID, remoteStorageURL string, user m.AuthUser, exposed, browseable, verbose bool) {
	checkRepoId(repoID)
	if remoteStorageURL == "" {
		log.Fatal("remoteStorageURL is a required parameter for creating a proxy repository")
	}
	createProxyRepo(nexusURL, repoID, "maven2", "release", remoteStorageURL, user, exposed, browseable, verbose)
}

func CreateNugetGroupRepo(nexusURL, repoID string, user m.AuthUser, verbose bool) {
	checkRepoId(repoID)
	createGroupRepo(nexusURL, repoID, "group", "nuget-group", user, verbose)
}

func createHostedRepo(nexusURL, repoID, provider, repoPolicy, writePolicy string, user m.AuthUser, verbose bool) {

	url := fmt.Sprintf("%s/service/local/repositories", nexusURL)

	repository := m.HostedRepository{
		Data: m.HostedRepositoryData{
			ID:               repoID,
			Name:             repoID,
			RepoType:         "hosted",
			RepoPolicy:       strings.ToUpper(repoPolicy),
			Provider:         provider,
			ProviderRole:     "org.sonatype.nexus.proxy.repository.Repository",
			Browseable:       true,
			Exposed:          true,
			WritePolicy:      writePolicy,
			ChecksumPolicy:   "IGNORE",
			Indexable:        true,
			NotFoundCacheTTL: 1440,
		},
	}

	body, err := json.Marshal(repository)
	u.Error(err, "Error creating the request body")

	req := u.CreateBaseRequest("POST", url, body, user, verbose)
	_, status := u.HTTPRequest(user, req, verbose)

	handleCreateStatus(status, repoID, repository.Data.RepoType)
}

func createProxyRepo(nexusURL, repoID, provider, repoPolicy, remoteStorageUrl string, user m.AuthUser, exposed, browseable, verbose bool) {

	url := fmt.Sprintf("%s/service/local/repositories", nexusURL)

	remoteStorage := m.ProxyRemoteStorage{
		RemoteStorageURL: remoteStorageUrl,
	}

	repository := m.ProxyRepository{
		Data: m.ProxyRepositoryData{
			ID:                    repoID,
			Name:                  repoID,
			RepoType:              "proxy",
			RepoPolicy:            strings.ToUpper(repoPolicy),
			Provider:              provider,
			ProviderRole:          "org.sonatype.nexus.proxy.repository.Repository",
			Browseable:            browseable,
			Exposed:               exposed,
			ChecksumPolicy:        "WARN",
			Indexable:             true,
			NotFoundCacheTTL:      1440,
			DownloadRemoteIndexes: true,
			ArtifactMaxAge:        -1,
			AutoBlockActive:       true,
			FileTypeValidation:    true,
			ItemMaxAge:            1440,
			MetadataMaxAge:        1440,
			RemoteStorage:         remoteStorage,
		},
	}

	body, err := json.Marshal(repository)
	u.Error(err, "Error creating the request body")

	req := u.CreateBaseRequest("POST", url, body, user, verbose)
	_, status := u.HTTPRequest(user, req, verbose)

	handleCreateStatus(status, repoID, repository.Data.RepoType)
}

func createGroupRepo(nexusURL, repoID, repoType, provider string, user m.AuthUser, verbose bool) {
	url := fmt.Sprintf("%s/service/local/repo_groups", nexusURL)
	repository := m.GroupRepository{
		Data: m.GroupRepositoryData{
			ID:       repoID,
			Name:     repoID,
			Provider: provider,
			Exposed:  true,
			Repositories:[]m.Repositories{},
		},
	}

	body, err := json.Marshal(repository)
	u.Error(err, "Error creating the request body")

	req := u.CreateBaseRequest("POST", url, body, user, verbose)
	_, status := u.HTTPRequest(user, req, verbose)

	handleCreateStatus(status, repoID, repoType)
}

func handleCreateStatus(status, repoID, repoType string) {
	switch status {
	case "201 Created":
		log.Printf("%s repository with ID=%s is created.\n", strings.Title(repoType), repoID)
	case "400 Bad Request":
		log.Printf("Repository with ID=%s already exists!\n", repoID)
	case "401 Unauthorized":
		log.Println("User could not be authenticated")
		os.Exit(1)
	default:
		panic(fmt.Sprintf("ERROR: call status=%v\n", status))
	}
}
