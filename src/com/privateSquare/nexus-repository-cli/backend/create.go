package backend

import (
	"com/privateSquare/nexus-repository-cli/model"
	"com/privateSquare/nexus-repository-cli/utils"
	"encoding/json"
	"log"
	"strings"
)

func CreateHostedRepo(user model.User, nexusUrl, repoId, repoType, repoPolicy, provider string, exposed, verbose bool) {

	url := nexusUrl + "/service/local/repositories"

	var writePolicy string
	if repoPolicy == "release" || repoPolicy == "mixed" {
		writePolicy = "ALLOW_WRITE_ONCE"
	} else if repoPolicy == "snapshot" {
		writePolicy = "ALLOW_WRITE"
	}

	repository := model.HostedRepository{
		Data: model.HostedRepositoryData{
			ID:               repoId,
			Name:             repoId,
			RepoType:         repoType,
			RepoPolicy:       strings.ToUpper(repoPolicy),
			Provider:         provider,
			ProviderRole:     "org.sonatype.nexus.proxy.repository.Repository",
			Browseable:       true,
			Exposed:          exposed,
			WritePolicy:      writePolicy,
			ChecksumPolicy:   "IGNORE",
			Indexable:        true,
			NotFoundCacheTTL: 1440,
		},
	}

	body, err := json.Marshal(repository)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, status := utils.HttpRequest(url, "POST", body, user.Username, user.Password, verbose)
	utils.PrintCreateStatus(status, repository.Data.ID, repository.Data.RepoType)
}

func CreateProxyRepo( user model.User, nexusUrl, repoId, repoType, repoPolicy, remoteStorageUrl, provider string, exposed, verbose bool) {

	url := nexusUrl + "/service/local/repositories"

	remoteStorage := model.ProxyRemoteStorage{
		RemoteStorageURL: remoteStorageUrl,
	}

	repository := model.ProxyRepository{
		Data: model.ProxyRepositoryData{
			ID:                    repoId,
			Name:                  repoId,
			RepoType:              repoType,
			RepoPolicy:            strings.ToUpper(repoPolicy),
			Provider:              provider,
			ProviderRole:          "org.sonatype.nexus.proxy.repository.Repository",
			Browseable:            true,
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
	if err != nil {
		log.Fatal(err)
		return
	}

	_, status := utils.HttpRequest(url, "POST", body, user.Username, user.Password, verbose)
	utils.PrintCreateStatus(status, repository.Data.ID, repository.Data.RepoType)
}

func CreateGroupRepo(user model.User, nexusUrl, repoId, repoType, repositories, provider string, exposed, verbose bool) {

	url := nexusUrl + "/service/local/repo_groups"

	var repositoriesModel []model.Repositories

	repositoriesArray := strings.Split(repositories, ",")

	for _, repository := range repositoriesArray{
		if CheckRepoExist(user, nexusUrl, repository, verbose){
			repositoriesModel = append(repositoriesModel, model.Repositories{ID: repository,})
		}else{
			log.Printf("Repository with ID=%s does not exist in Nexus, hence not adding it to the group repository", repository)
		}

	}
	repository := model.GroupRepository{
		Data: model.GroupRepositoryData{
			ID:           repoId,
			Name:         repoId,
			Provider:     provider,
			Exposed:      exposed,
			Repositories: repositoriesModel,
		},
	}

	body, err := json.Marshal(repository)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, status := utils.HttpRequest(url, "POST", body, user.Username, user.Password, verbose)
	utils.PrintCreateStatus(status, repository.Data.ID, repoType)
}
