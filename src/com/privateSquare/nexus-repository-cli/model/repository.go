package model

// RepositoryData represents the Nexus repository data
type RepositoryData struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	RepoType   string `json:"repoType"`
	RepoPolicy string `json:"repoPolicy"`
	Provider   string `json:"provider"`
	Format     string `json:"format"`
}

// Repository represents a Nexus repository
type Repository struct {
	Data []RepositoryData `json:"data"`
}

// HostedRepositoryData represents the Nexus hosted repository data
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

// HostedRepository represents a Nexus hosted repository
type HostedRepository struct {
	Data HostedRepositoryData `json:"data"`
}

// ProxyRemoteStorage represents the remoteStorageUrl for a proxy repository
type ProxyRemoteStorage struct {
	RemoteStorageURL string `json:"remoteStorageUrl"`
}

// ProxyRepositoryData represents the Nexus proxy repository data
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

// ProxyRepository represents a Nexus proxy repository
type ProxyRepository struct {
	Data ProxyRepositoryData `json:"data"`
}

// Repositories represents the repositores for a group repository
type Repositories struct {
	ID string `json:"id"`
}

// GroupRepositoryData represents the Nexus group repository data
type GroupRepositoryData struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Provider     string         `json:"provider"`
	Exposed      bool           `json:"exposed"`
	Repositories []Repositories `json:"repositories"`
}

// GroupRepository represents a Nexus group repository
type GroupRepository struct {
	Data GroupRepositoryData `json:"data"`
}
