# Nexus Repository CLI

*_This only works on nexus 2_*

## Usage:

```go
Usage: ./nexus-repository-cli.exe [option] [parameters...]

 [options]
  -list
        List the repositories in Nexus. Optional parameters: repoType, repoPolicy
  -create
        Create a repository in Nexus. Required parameter: repoId, repoType, provider, repoPolicy (only for maven2). Optional parameter: exposed
  -delete
        Delete a repository in Nexus. Required parameter: repoId
  -addRepoToGroup
        Add a reposirory to a group repository. Required paramters: repoId, repositories

 [parameters]
  -nexusUrl string
        Nexus server URL (default "http://localhost:8081/nexus")
  -exposed
        Set this flag to expose the repository in nexus.
  -username string
        Username for authentication
  -password string
        Password for authentication
  -repoId string
        ID of the Repository
  -repoType string
        Type of a repository. Possible values : hosted/proxy/group
  -repoPolicy string
        Policy of the hosted repository. Possible values : snapshot/release
  -provider string
        Repository provider. Possible values: maven2/npm/nuget
  -remoteStorageUrl string
        Remote storage url to proxy in Nexus
  -repositories string
        Comma separated value of repositories to be added to a group.
  -verbose
        Set this flag for Debug logs.
```

## Examples:

### [option] -list

```sh
./nexus-repository-cli.exe -username ***** -password ***** -list # Prints all repositories
./nexus-repository-cli.exe -username ***** -password ***** -list -repoType hosted # Prints all the hosted repositories
./nexus-repository-cli.exe -username ***** -password ***** -list -provider maven2 # Prints all the maven repositories
./nexus-repository-cli.exe -username ***** -password ***** -list -repoType hosted -provider maven2 # Prints all the hosted maven repositories
./nexus-repository-cli.exe -username ***** -password ***** -list -repoType hosted -provider maven2 -repoPolicy release # Prints all the hosted maven Release repositories
```

### [option] -create

#### Maven repository

```sh
./nexus-repository-cli.exe -username ***** -password ***** -create -repoId ATS-maven-snapshot -repoType hosted -provider maven -repoPolicy snapshot -exposed
./nexus-repository-cli.exe -username ***** -password ***** -create -repoId ATS-maven-releases -repoType hosted -provider maven -repoPolicy release -exposed
./nexus-repository-cli.exe -username ***** -password ***** -create -repoId ATS-maven-proxy -repoType proxy -provider maven -remoteStorageUrl https://repo1.maven.org/maven2/
./nexus-repository-cli.exe -username ***** -password ***** -create -repoId ATS-maven-group -repoType group -provider maven -repositories ATS-maven-snapshot,ATS-maven-releases,ATS-maven-proxy -exposed
```

#### NPM repository

```sh
./nexus-repository-cli.exe -username ***** -password ***** -create -repoId ATS-npm-releases -repoType hosted -provider npm -exposed
./nexus-repository-cli.exe -username ***** -password ***** -create -repoId ATS-npm-proxy -repoType proxy -provider npm -remoteStorageUrl http://registry.npmjs.org/
./nexus-repository-cli.exe -username ***** -password ***** -create -repoId ATS-npm-group -repoType group -provider npm -repositories ATS-npm-releases,ATS-npm-proxy -exposed
```

#### Nuget repository

```sh
./nexus-repository-cli.exe -username ***** -password ***** -create -repoId ATS-nuget-gallery -repoType hosted -provider nuget -exposed
./nexus-repository-cli.exe -username ***** -password ***** -create -repoId ATS-nuget-proxy -repoType proxy -provider nuget -remoteStorageUrl https://www.nuget.org/api/v2/
./nexus-repository-cli.exe -username ***** -password ***** -create -repoId ATS-nuget-group -repoType group -provider nuget -repositories ATS-nuget-gallery,ATS-nuget-proxy -exposed
```

### [option] -delete

```sh
./nexus-repository-cli.exe -username ***** -password ***** -delete -repoId ATS-maven-snapshot
```