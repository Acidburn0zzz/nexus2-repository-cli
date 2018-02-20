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

type PrivilegesData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
}

type Privileges struct {
	Data []PrivilegesData `json:"data"`
}

type PrivilegeCreateData struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	RepositoryID       string `json:"repositoryId"`
	RepositoryGroupID  string `json:"repositoryGroupId"`
	Type               string `json:"type"`
	RepositoryTargetID string `json:"repositoryTargetId"`
	Method 			   []string `json:"method"`
}

type PrivilegesCreate struct {
	Data PrivilegeCreateData `json:"data"`
}

// roles

type RoleData struct {
	ID             string        `json:"id"`
	Name           string        `json:"name"`
	Description    string        `json:"description"`
	SessionTimeout int           `json:"sessionTimeout"`
	Roles          []string 	 `json:"roles"`
	Privileges     []string 	 `json:"privileges"`
}

type Role struct {
	Data RoleData `json:"data"`
}

type Roles struct {
	Data []RoleData `json:"data"`
}