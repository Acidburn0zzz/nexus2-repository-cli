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

func CreateMavenHostedRepo(nexusURL, repoId string, user m.AuthUser, release, verbose bool) {
	CheckRepoId(repoId)
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
	createHostedRepo(nexusURL, repoId, "maven2", repoPolicy, writePolicy, user, verbose)
}

func CreateMavenProxyRepo(nexusURL, repoId, remoteStorageURL string, user m.AuthUser, exposed, browseable, verbose bool) {
	CheckRepoId(repoId)
	if remoteStorageURL == "" {
		log.Fatal("remoteStorageURL is a required parameter for creating a proxy repository")
	}
	createProxyRepo(nexusURL, repoId, "maven2", "release", remoteStorageURL, user, exposed, browseable, verbose)
}

func CreateMavenGroupRepo(nexusURL, repoId string, user m.AuthUser, verbose bool) {
	CheckRepoId(repoId)
	createGroupRepo(nexusURL, repoId, "group", "maven2", user, verbose)
}

func CreateNPMHostedRepo(nexusURL, repoId string, user m.AuthUser, verbose bool) {
	CheckRepoId(repoId)
	createHostedRepo(nexusURL, repoId, "npm-hosted", "mixed", "ALLOW_WRITE_ONCE", user, verbose)
}

func CreateNPMProxyRepo(nexusURL, repoId, remoteStorageURL string, user m.AuthUser, exposed, browseable, verbose bool) {
	CheckRepoId(repoId)
	if remoteStorageURL == "" {
		log.Fatal("remoteStorageURL is a required parameter for creating a proxy repository")
	}
	createProxyRepo(nexusURL, repoId, "maven2", "release", remoteStorageURL, user, exposed, browseable, verbose)
}

func CreateNPMGroupRepo(nexusURL, repoId string, user m.AuthUser, verbose bool) {
	CheckRepoId(repoId)
	createGroupRepo(nexusURL, repoId, "group", "npm-group", user, verbose)
}

func CreateNugetHostedRepo(nexusURL, repoId string, user m.AuthUser, verbose bool) {
	CheckRepoId(repoId)
	createHostedRepo(nexusURL, repoId, "nuget-proxy", "mixed", "ALLOW_WRITE_ONCE", user, verbose)
}

func CreateNugetProxyRepo(nexusURL, repoId, remoteStorageURL string, user m.AuthUser, exposed, browseable, verbose bool) {
	CheckRepoId(repoId)
	if remoteStorageURL == "" {
		log.Fatal("remoteStorageURL is a required parameter for creating a proxy repository")
	}
	createProxyRepo(nexusURL, repoId, "maven2", "release", remoteStorageURL, user, exposed, browseable, verbose)
}

func CreateNugetGroupRepo(nexusURL, repoId string, user m.AuthUser, verbose bool) {
	CheckRepoId(repoId)
	createGroupRepo(nexusURL, repoId, "group", "nuget-group", user, verbose)
}

func createHostedRepo(nexusURL, repoId, provider, repoPolicy, writePolicy string, user m.AuthUser, verbose bool) {

	url := fmt.Sprintf("%s/service/local/repositories", nexusURL)

	repository := m.HostedRepository{
		Data: m.HostedRepositoryData{
			ID:               repoId,
			Name:             repoId,
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

	handleCreateStatus(status, repoId, repository.Data.RepoType)
}

func createProxyRepo(nexusURL, repoId, provider, repoPolicy, remoteStorageUrl string, user m.AuthUser, exposed, browseable, verbose bool) {

	url := fmt.Sprintf("%s/service/local/repositories", nexusURL)

	remoteStorage := m.ProxyRemoteStorage{
		RemoteStorageURL: remoteStorageUrl,
	}

	repository := m.ProxyRepository{
		Data: m.ProxyRepositoryData{
			ID:                    repoId,
			Name:                  repoId,
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

	handleCreateStatus(status, repoId, repository.Data.RepoType)
}

func createGroupRepo(nexusURL, repoId, repoType, provider string, user m.AuthUser, verbose bool) {

	url := fmt.Sprintf("%s/service/local/repo_groups", nexusURL)

	repository := m.GroupRepository{
		Data: m.GroupRepositoryData{
			ID:       repoId,
			Name:     repoId,
			Provider: provider,
			Exposed:  true,
		},
	}

	body, err := json.Marshal(repository)
	u.Error(err, "Error creating the request body")

	req := u.CreateBaseRequest("POST", url, body, user, verbose)
	_, status := u.HTTPRequest(user, req, verbose)

	handleCreateStatus(status, repoId, repoType)
}

func handleCreateStatus(status, repoId, repoType string) {
	switch status {
	case "201 Created":
		log.Printf("%s repository with ID=%s is created.\n", strings.Title(repoType), repoId)
	case "400 Bad Request":
		log.Printf("Repository with ID=%s already exists!\n", repoId)
	case "401 Unauthorized":
		log.Println("User could not be authenticated")
		os.Exit(1)
	default:
		panic(fmt.Sprintf("ERROR: call status=%v\n", status))
	}
}
