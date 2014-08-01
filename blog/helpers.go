package blog

import "log"

// Handles errors and outputs custom message, followed
// by error string
func HandleError(err error, msg string) {
	if err != nil {
		log.Fatal(msg, " (", err, ")")
	}
}
