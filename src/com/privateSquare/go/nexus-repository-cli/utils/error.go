package utils

import "log"

/*
Error prints error
@param err error  error details
@return void
*/
func Error(err error, errorMessage string) {
	if err != nil {
		log.Println(errorMessage)
		log.Fatal(err)
	}
}
