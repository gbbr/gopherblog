package helpers

import "fmt"

// Handles errors, and outputs custom message, followed
// by error string
func HandleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg, "(", err, ")")
	}
}
