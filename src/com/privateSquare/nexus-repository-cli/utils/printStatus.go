package utils

import (
	"fmt"
	"log"
	"strings"
)

// PrintCreateStatus prints the status when a repository create us invoked based on the status of the response
func PrintCreateStatus(status, repoId, repoType string) {
	switch status {
	case "201 Created":
		log.Printf("%s repository with ID=%s is created.\n", strings.Title(repoType), repoId)
	case "400 Bad Request":
		log.Printf("Repository with ID=%s already exists!\n", repoId)
	case "401 Unauthorized":
		log.Printf("User could not be authenticated")
	default:
		panic(fmt.Sprintf("ERROR: call status=%v\n", status))
	}
}
