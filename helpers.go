package bam

import (
	"fmt"
	"log"

	multierror "github.com/hashicorp/go-multierror"
)

// LogoutClientIfError will log out the client session if an error is passed in
func LogoutClientIfError(client ProteusAPI, err error, msg string) error {
	if err != nil {
		var result error
		result = multierror.Append(err)

		if lerr := client.Logout(); lerr != nil {
			result = multierror.Append(lerr)
		}
		log.Printf("[INFO] BlueCat Logout was successful")
		return fmt.Errorf(msg, result)
	}
	return nil
}
