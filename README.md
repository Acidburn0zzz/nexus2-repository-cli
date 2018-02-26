# Nexus Repository CLI

*_The cli only works with nexus 2 repository manager_*

## Usage

```console
Usage of ./nexus2-repository-cli.exe:
  -addRepoToGroup
        Add a reposirory to a group repository. Required parameters: repoID, repositories.
  -appSysLetters string
        Applictaion system letters of an application. (3 or 4 letter code)
  -browseable
        Set this flag to make the repository browseable in nexus.
  -createMavenGroupRepo
        Create a maven group repository. Required parameters: repoID.
  -createMavenHostedRepo
        Create a maven hosted repository (By default a snapshot repository is created). Required parameters: repoID Optional parameter: release (creates a release repository).
  -createMavenProxyRepo
        Create a maven proxy repository. Required parameters: repoID, remoteStorageURL. Optional parameters: exposed, browseable.
  -createMavenTarget
        Create a maven repository target. Required parameters: repoTargetName, patternExpression.
  -createPrivileges
        Create repository privileges. Required parameters: privilegeName, repoID, repoTargetName.
  -createRole
        Create roles. Required parameters: roleName, privileges, roles.
  -deleteGroupRepo
        Deletes a group repository. Required parameter: repoID.
  -deletePrivileges
        Delete repository privileges. Required parameters: privilegeName.
  -deleteRepo
        Deletes a hosted/proxy repository. Required parameter: repoID.
  -deleteRole
        Delete roles. Required parameters: roleName.
  -deleteTarget
        Delete a repository target. Required parameters: repoTargetName.
  -exposed
        Set this flag to expose the repository in nexus.
  -javaOffboarding
        Remove an application from the java repo structure. Required paramters: appSysLetters.
  -javaOnboarding
        Create a new space for a application in the java repo structure. Required paramters: appSysLetters.
  -list
        List the repositories in Nexus. Optional parameters: repoType, repoPolicy
  -nexusURL string
        Nexus server URL. (default "http://localhost:8081/nexus")
  -password string
        Password for authentication.
  -pattern string
        Repository target pattern expression. Can be comma separated values.
  -privilegeName string
        Repository Privilege name.
  -privileges string
        Comma separated privilege name values.
  -provider string
        Repository provider. Possible values: maven2/npm/nuget.
  -release
        Set this flag for creating a maven release repository.
  -remoteStorageURL string
        Remote storage url to proxy in Nexus.
  -repoID string
        ID of a Repository.
  -repoPolicy string
        Policy of the hosted repository. Possible values : snapshot/release.
  -repoTargetName string
        Repository target name.
  -repoType string
        Type of a repository. Possible values : hosted/proxy/group.
  -repositories string
        Comma separated value of repositories to be added to a group.
  -roleName string
        Role name.
  -roles string
        Comma separated role name values.
  -username string
        Username for authentication.
  -verbose
        Set this flag for Debug logs.
```

## Print Help

```console
./nexus2-repository-cli.exe -help
```

## Examples

### -list

```sh
./nexus-repository-cli.exe -username ***** -password ***** -list # Prints all repositories
./nexus-repository-cli.exe -username ***** -password ***** -list -repoType hosted # Prints all the hosted repositories
./nexus-repository-cli.exe -username ***** -password ***** -list -provider maven2 # Prints all the maven repositories
./nexus-repository-cli.exe -username ***** -password ***** -list -repoType hosted -provider maven2 # Prints all the hosted maven repositories
# TODO search on -repoPolicy is not working
./nexus-repository-cli.exe -username ***** -password ***** -list -repoType hosted -provider maven2 -repoPolicy release # Prints all the hosted maven Release repositories
```