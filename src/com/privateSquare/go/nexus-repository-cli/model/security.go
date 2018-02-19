package model

//targets

type RepositoryTargetData struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	ContentClass string   `json:"contentClass"`
	Patterns     []string `json:"patterns"`
}

type RepositoryTargets struct {
	Data []RepositoryTargetData `json:"data"`
}

type RepositoryTargetCreateData struct {
	Name         string   `json:"name"`
	ContentClass string   `json:"contentClass"`
	Patterns     []string `json:"patterns"`
}

type RepositoryTargetCreate struct {
	Data RepositoryTargetCreateData `json:"data"`
}

//privileges

type RepoPrivilegesData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
}

type RepoPrivileges struct {
	Data []RepoPrivilegesData `json:"data"`
}

type RepoPrivilegeCreateData struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	RepositoryID       string `json:"repositoryId"`
	RepositoryGroupID  string `json:"repositoryGroupId"`
	Type               string `json:"type"`
	RepositoryTargetID string `json:"repositoryTargetId"`
	Method 			   []string `json:"method"`
}

type RepoPrivilegesCreate struct {
	Data RepoPrivilegeCreateData `json:"data"`
}

