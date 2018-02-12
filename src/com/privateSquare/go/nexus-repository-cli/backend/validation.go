package backend

import (
	"log"
)

// CheckRepoType checks is a valid repoType is entered
func CheckRepoType(repoType string) {
	if repoType != "" && repoType != "hosted" && repoType != "group" && repoType != "proxy" {
		log.Fatal("repoType value is invalid. Possible values are hosted/group/proxy")
	}
}

// CheckProvider checks if a valid provider is entered
func CheckProvider(provider string) {
	if provider != "" && provider != "maven2" && provider != "npm" && provider != "nuget" {
		log.Fatal("provider value is invalid. Possible values are maven2/npm/nuget")
	}
}

// CheckMavenRepoPolicy checks if a valid repoPolicy is entered
func CheckMavenRepoPolicy(repoPolicy string) {
	if repoPolicy != "" && repoPolicy != "release" && repoPolicy != "snapshot" {
		log.Fatal("repoPolicy value is invalid. Possible values are snapshot/release")
	}
}
