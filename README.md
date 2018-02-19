# Nexus Repository CLI

*_This only works with nexus 2 repository manager_*

## Usage

```go
Usage of nexus2-repository-cli.exe:
  -addRepoToGroup
        Add a reposirory to a group repository. Required parameters: repoId, repositories.
  -browseable
        Set this flag to make the repository browseable in nexus.
  -createMavenGroupRepo
        Create a maven group repository. Required parameters: repoId.
  -createMavenHostedRepo
        Create a maven hosted repository (By default a snapshot repository is created). Required parameters: repoId Optional parameter: release (creates a release repository).
  -createMavenProxyRepo
        Create a maven proxy repository. Required parameters: repoId, remoteStorageURL. Optional parameters: exposed, browseable.
  -createMavenTarget
        Create a maven repository target. Required parameters: repoTargetName, patternExpression.
  -createPrivileges
        Create repository privileges. Required parameters: repoPrivilegeName, repoTargetName.
  -delete
        Delete a repository in Nexus. Required parameter: repoId.
  -deletePrivileges
        Delete repository privileges. Required parameters: repoPrivilegeName.
  -deleteTarget
        Delete a repository target. Required parameters: repoTargetName.
  -exposed
        Set this flag to expose the repository in nexus.
  -list
        List the repositories in Nexus. Optional parameters: repoType, repoPolicy
  -nexusURL string
        Nexus server URL. (default "http://localhost:8081/nexus")
  -password string
        Password for authentication.
  -patternExpression string
        Repository target pattern expression. Can be comma separated values.
  -provider string
        Repository provider. Possible values: maven2/npm/nuget.
  -release
        Set this flag for creating a maven release repository.
  -remoteStorageURL string
        Remote storage url to proxy in Nexus.
  -repoId string
        ID of a Repository.
  -repoPolicy string
        Policy of the hosted repository. Possible values : snapshot/release.
  -repoPrivilegeName string
        Repository Privilege name.
  -repoTargetName string
        Repository target name.
  -repoType string
        Type of a repository. Possible values : hosted/proxy/group.
  -username string
        Username for authentication.
  -verbose
        Set this flag for Debug logs.
```

## Examples

### [option] -list

```sh
./nexus-repository-cli.exe -username ***** -password ***** -list # Prints all repositories
./nexus-repository-cli.exe -username ***** -password ***** -list -repoType hosted # Prints all the hosted repositories
./nexus-repository-cli.exe -username ***** -password ***** -list -provider maven2 # Prints all the maven repositories
./nexus-repository-cli.exe -username ***** -password ***** -list -repoType hosted -provider maven2 # Prints all the hosted maven repositories
# TODO search on -repoPolicy is not working
./nexus-repository-cli.exe -username ***** -password ***** -list -repoType hosted -provider maven2 -repoPolicy release # Prints all the hosted maven Release repositories
```