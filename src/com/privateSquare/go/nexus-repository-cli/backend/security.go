package backend

import (
	"fmt"
	u "com/privateSquare/go/nexus-repository-cli/utils"
	m "com/privateSquare/go/nexus-repository-cli/model"
	"encoding/json"
	"log"
	"os"
	"strings"
	"regexp"
)

func CreateMavenRepoTarget(nexusURL, repoTargetName, patternExpressions string, user m.AuthUser, verbose bool){
	if repoTargetName == "" || patternExpressions == "" {
		log.Println("repoTargetName and patternExpressions are required parameters for creating a repository target.")
		os.Exit(1)
	}
	createRepoTarget(nexusURL, repoTargetName, patternExpressions, "maven2", user, verbose)
}

func createRepoTarget(nexusURL, repoTargetName, patternExpressions, contentClass string ,user m.AuthUser, verbose bool){
	if !repoTargetExists(nexusURL, repoTargetName, user, verbose){
		url := fmt.Sprintf("%s/service/local/repo_targets", nexusURL)
		patternExpressionsArray := strings.Split(patternExpressions, ",")
		repoTarget := m.RepositoryTargetCreate{
			Data: m.RepositoryTargetCreateData{
				Name:repoTargetName,
				ContentClass:contentClass,
				Patterns: patternExpressionsArray,
			},
		}
		body, err := json.Marshal(repoTarget)
		u.Error(err, "Error creating the request body")
		req := u.CreateBaseRequest("POST", url, body, user, verbose)
		_, status := u.HTTPRequest(user, req, verbose)
		if status == "201 Created"{
			log.Printf("Repository target '%s' is created.", repoTargetName)
		}else{
			log.Printf("Repository target create request failed with status : %s\n.", status)
			log.Printf("Activate -verbose flag for more details on the error.\n", status)
			os.Exit(1)
		}
	} else {
		log.Printf("Repository target '%s' already exists.", repoTargetName)
		os.Exit(1)
	}
}

func DeleteRepoTarget(nexusURL, repoTargetName string ,user m.AuthUser, verbose bool){
	if repoTargetName == "" {
		log.Println("repoTargetName is a required paramteter for deleting a repository target.")
		os.Exit(1)
	}
	if repoTargetExists(nexusURL, repoTargetName, user, verbose){
		repoTargetID := getRepoTargetID(nexusURL, repoTargetName, user, verbose)
		url := fmt.Sprintf("%s/service/local/repo_targets/%s", nexusURL, repoTargetID)
		req := u.CreateBaseRequest("DELETE", url, nil, user, verbose)
		u.HTTPRequest(user, req, verbose)
		log.Printf("Repository target '%s' deleted.\n", repoTargetName)
	}else {
		log.Printf("Repository target '%s' does not exists.", repoTargetName)
		os.Exit(1)
	}
}

func getRepoTargetID (nexusURL, repoTargetName string, user m.AuthUser, verbose bool) string{
	url := fmt.Sprintf("%s/service/local/repo_targets", nexusURL)
	req := u.CreateBaseRequest("GET", url, nil, user, verbose)
	respBody, _ := u.HTTPRequest(user, req, verbose)
	var (
		repoTargetId string
		repoTargets m.RepositoryTargets
	)
	json.Unmarshal(respBody, &repoTargets)
	for _, repoTarget := range repoTargets.Data{
		if repoTarget.Name == repoTargetName {
			repoTargetId = repoTarget.ID
		}
	}
	return repoTargetId
}

func repoTargetExists (nexusURL, repoTargetName string, user m.AuthUser, verbose bool) bool{
	var isExists bool
	repoTargetID := getRepoTargetID(nexusURL, repoTargetName, user, verbose)
	if repoTargetID != ""{
		isExists = true
	} else {
		isExists = false
	}
	return isExists
}

func CreatePrivileges(nexusURL, privilegeName, repoTargetName string ,user m.AuthUser, verbose bool){
	if privilegeName == "" || repoTargetName == "" {
		log.Println("privilegeName and repoTargetName are required parameters for creating repository privileges.")
		os.Exit(1)
	}
	if !repoPrivilegesExists(nexusURL, privilegeName, user, verbose){
		url := fmt.Sprintf("%s/service/local/privileges_target", nexusURL)
		privileges := m.PrivilegesCreate{
			Data: m.PrivilegeCreateData{
				Name:privilegeName,
				Description:privilegeName,
				RepositoryID:"",
				RepositoryGroupID:"",
				Type:"target",
				RepositoryTargetID:getRepoTargetID(nexusURL, repoTargetName, user, verbose),
				Method: []string{"create", "read", "update", "delete"},
			},
		}
		body, err := json.Marshal(privileges)
		u.Error(err, "Error creating the request body")
		req := u.CreateBaseRequest("POST", url, body, user, verbose)
		u.HTTPRequest(user, req, verbose)
		log.Printf("CRUD Repository privileges for '%s' created.\n", privilegeName)
	}else {
		log.Printf("Repository privileges for '%s' already exists.", privilegeName)
		os.Exit(1)
	}
}

func DeletePrivileges(nexusURL, privilegeName string ,user m.AuthUser, verbose bool){
	if privilegeName == "" {
		log.Println("privilegeName is a required parameter for deleting repository privileges.")
		os.Exit(1)
	}
	if repoPrivilegesExists(nexusURL, privilegeName, user, verbose){
		repoPrivilegesID := getPrivilegesID(nexusURL, privilegeName, user, verbose)
		for _, repoPrivilegeID := range repoPrivilegesID{
			url := fmt.Sprintf("%s/service/local/privileges/%s", nexusURL, repoPrivilegeID)
			req := u.CreateBaseRequest("DELETE", url, nil, user, verbose)
			u.HTTPRequest(user, req, verbose)
		}
		log.Printf("All Repository privileges for '%s' deleted.\n", privilegeName)
	}else{
		log.Printf("Repository privileges for '%s' does not exists.", privilegeName)
		os.Exit(1)
	}
}

func getPrivilegesID (nexusURL, privilegeName string, user m.AuthUser, verbose bool) []string{
	url := fmt.Sprintf("%s/service/local/privileges", nexusURL)
	req := u.CreateBaseRequest("GET", url, nil, user, verbose)
	respBody, _ := u.HTTPRequest(user, req, verbose)
	var (
		privileges m.Privileges
		privilegesID []string
	)
	json.Unmarshal(respBody, &privileges)
	cRegexp, err := regexp.Compile(privilegeName + " ")
	u.Error(err, "There was a error compiling the regex.")
	for _, privilege := range privileges.Data{
		if cRegexp.MatchString(privilege.Name) {
			privilegesID = append(privilegesID, privilege.ID)
		}
	}
	return privilegesID
}

func getPrivilegeID (nexusURL, privilegeName string, user m.AuthUser, verbose bool) string{
	url := fmt.Sprintf("%s/service/local/privileges", nexusURL)
	req := u.CreateBaseRequest("GET", url, nil, user, verbose)
	respBody, _ := u.HTTPRequest(user, req, verbose)
	var (
		privileges m.Privileges
		privilegesID string
	)
	json.Unmarshal(respBody, &privileges)
	for _, repoPrivilege := range privileges.Data{
		if repoPrivilege.Name == privilegeName {
			privilegesID = repoPrivilege.ID
		}
	}
	return privilegesID
}

func repoPrivilegesExists (nexusURL, repoPrivilegeName string, user m.AuthUser, verbose bool) bool{
	var isExists bool
	repoPrivileges := getPrivilegesID(nexusURL, repoPrivilegeName, user, verbose)
	if len(repoPrivileges) > 0 {
		isExists = true
	}else {
		isExists = false
	}
	return isExists
}

func privilegeExists (nexusURL, privilegeName string, user m.AuthUser, verbose bool) bool{
	var isExists bool
	privilegeID := getPrivilegeID(nexusURL,privilegeName, user, verbose)
	if privilegeID != "" {
		isExists = true
	}else {
		isExists = false
	}
	return isExists
}

func CreateRole (nexusURL, roleName string, privileges, roles string, user m.AuthUser, verbose bool){
	if roleName == "" || privileges == "" || roles == "" {
		log.Println("roleName, privileges and roles are required paramters for creating a role.")
		os.Exit(1)
	}
	var (
		previlegesIDList []string
		rolesIDList [] string
	)
	if !roleExists(nexusURL,roleName,user,verbose){
		privilegesList := strings.Split(privileges, ",")
		for _, privilege := range privilegesList{
			if privilegeExists(nexusURL, privilege, user, verbose){
				previlegesIDList = append(previlegesIDList, getPrivilegeID(nexusURL, privilege, user, verbose))
			} else {
				log.Printf("Privilege '%s' does not exist, hence not adding the prvilege to the role '%s'.\n", privilege, roleName)
			}
		}
		rolesList := strings.Split(roles, ",")
		for _, role := range rolesList{
			if roleExists(nexusURL, role, user, verbose){
				rolesIDList = append(rolesIDList, getRoleID(nexusURL, role, user, verbose))
			}else {
				log.Printf("Role '%s' does not exist, hence not adding the role to the role '%s'.\n", role, roleName)
			}
		}
		if len(previlegesIDList) == 0 || len(rolesIDList) == 0{
			log.Printf("Need to add atlease one privilege/role for creating a new role.\n")
			os.Exit(1)
		}else{
			createRoleMapping(nexusURL, roleName, previlegesIDList, rolesList, user, verbose)
		}
	}else{
		log.Printf("Role '%s' already exists.\n", roleName)
		os.Exit(1)
	}

}

func createRoleMapping(nexusURL, roleName string, privilegesID, rolesID []string, user m.AuthUser, verbose bool){
	if !roleExists(nexusURL,roleName,user,verbose){
		url := fmt.Sprintf("%s/service/local/roles", nexusURL)
		rolesData := m.RoleData{
			ID:roleName,
			Name:roleName,
			Description:fmt.Sprintf("External mapping for %s (LDAP)", roleName),
			SessionTimeout:60,
			Privileges:privilegesID,
			Roles:rolesID,
		}
		roles := m.Role{Data:rolesData}
		body, err := json.Marshal(roles)
		u.Error(err, "Error creating the request body")
		req := u.CreateBaseRequest("POST", url, body, user, verbose)
		u.HTTPRequest(user, req, verbose)
		log.Printf("Role '%s' created.\n", roleName)
	}else {
		log.Printf("Role '%s' already exists.\n", roleName)
		os.Exit(1)
	}
}

func DeleteRole (nexusURL, roleName string, user m.AuthUser, verbose bool){
	if roleName == "" {
		log.Println("roleName is a required parameter for deleting a role.")
		os.Exit(1)
	}
	if roleExists(nexusURL, roleName, user, verbose){
		url := fmt.Sprintf("%s/service/local/roles/%s", nexusURL, roleName)
		req := u.CreateBaseRequest("DELETE", url, nil, user, verbose)
		u.HTTPRequest(user, req, verbose)
		log.Printf("Role '%s' deleted.\n", roleName)
	} else {
		log.Printf("Role '%s' does not exixts.\n", roleName)
		os.Exit(1)
	}
}

func AddPrivilegesToRole(){
	//TODO
}

func getRoleID(nexusURL, roleName string, user m.AuthUser, verbose bool) string{
	var (
		roleID string
		roles m.Roles
	)
	url := fmt.Sprintf("%s/service/local/roles", nexusURL)
	req := u.CreateBaseRequest("GET", url, nil, user, verbose)
	respBody, _ := u.HTTPRequest(user, req, verbose)
	json.Unmarshal(respBody, &roles)
	for _, role := range roles.Data{
		if roleName == role.Name{
			roleID = role.ID
		}
	}
	return roleID
}

func roleExists(nexusURL, roleName string, user m.AuthUser, verbose bool) bool {
	var isExists bool
	roleID := getRoleID(nexusURL, roleName, user, verbose)
	if roleID != "" {
		isExists = true
	}else {
		isExists = false
	}
	return isExists
}