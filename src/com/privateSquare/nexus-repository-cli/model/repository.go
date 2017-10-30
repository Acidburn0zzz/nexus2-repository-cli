package model

type RepositoryData struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	RepoType   string `json:"repoType"`
	RepoPolicy string `json:"repoPolicy"`
	Provider   string `json:"provider"`
	Format     string `json:"format"`
}

type Repository struct {
	Data []RepositoryData `json:"data"`
}

type HostedRepositoryData struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	RepoType         string `json:"repoType"`
	RepoPolicy       string `json:"repoPolicy"`
	Provider         string `json:"provider"`
	ProviderRole     string `json:"providerRole"`
	Browseable       bool   `json:"browseable"`
	Exposed          bool   `json:"exposed"`
	WritePolicy      string `json:"writePolicy"`
	ChecksumPolicy   string `json:"checksumPolicy"`
	Indexable        bool   `json:"indexable"`
	NotFoundCacheTTL int    `json:"notFoundCacheTTL"`
}

type HostedRepository struct {
	Data HostedRepositoryData `json:"data"`
}

type ProxyRemoteStorage struct {
	RemoteStorageURL string `json:"remoteStorageUrl"`
}

type ProxyRepositoryData struct {
	ID                    string             `json:"id"`
	Name                  string             `json:"name"`
	RepoType              string             `json:"repoType"`
	RepoPolicy            string             `json:"repoPolicy"`
	Provider              string             `json:"provider"`
	ProviderRole          string             `json:"providerRole"`
	Browseable            bool               `json:"browseable"`
	Exposed               bool               `json:"exposed"`
	ChecksumPolicy        string             `json:"checksumPolicy"`
	Indexable             bool               `json:"indexable"`
	NotFoundCacheTTL      int                `json:"notFoundCacheTTL"`
	DownloadRemoteIndexes bool               `json:"downloadRemoteIndexes"`
	ArtifactMaxAge        int                `json:"artifactMaxAge"`
	AutoBlockActive       bool               `json:"autoBlockActive"`
	FileTypeValidation    bool               `json:"fileTypeValidation"`
	ItemMaxAge            int                `json:"itemMaxAge"`
	MetadataMaxAge        int                `json:"metadataMaxAge"`
	RemoteStorage         ProxyRemoteStorage `json:"remoteStorage"`
}

type ProxyRepository struct {
	Data ProxyRepositoryData `json:"data"`
}

type Repositories struct {
	ID string `json:"id"`
}

type GroupRepositoryData struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Provider     string         `json:"provider"`
	Exposed      bool           `json:"exposed"`
	Repositories []Repositories `json:"repositories"`
}

type GroupRepository struct {
	Data GroupRepositoryData `json:"data"`
}