package bam

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fiorix/wsdl2go/soap"
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

// Client logs you in to BAM and keeps your session cookie
func Client(username string, password string, endpoint string) (ProteusAPI, error) {
	//var response *http.Response
	cli := soap.Client{
		URL:       "https://" + endpoint + "/Services/API?wsdl",
		Namespace: Namespace,
		Pre:       setBlueCatAuthToken,
		Post:      getBlueCatAuthToken,
	}
	soapService := NewProteusAPI(&cli)
	log.Printf("[INFO] BlueCat URL is: %s", cli.URL)
	err := soapService.Login(username, password)
	if err != nil {
		return nil, fmt.Errorf("Login error: %s", err)
	}
	log.Printf("[INFO] BlueCat Login was successful")

	return soapService, nil
}

func setBlueCatAuthToken(request *http.Request) {
	//a session cookie is required for all calls except Login
	for i := range sessionCookies {
		request.AddCookie(sessionCookies[i])
	}
}

func getBlueCatAuthToken(response *http.Response) {
	//response.Cookies() is usually empty except when calling Login
	//couldn't think of a better way to handle this
	sessionCookies = append(sessionCookies, response.Cookies()...)
}
