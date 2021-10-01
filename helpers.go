package gobam

import (
	"crypto/tls"
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

// LogoutClientWithError will log out the client session with the specified error message
func LogoutClientWithError(client ProteusAPI, msg string) error {
	var result error
	err := fmt.Errorf(msg)
	result = multierror.Append(err)

	if lerr := client.Logout(); lerr != nil {
		result = multierror.Append(lerr)
	}
	log.Printf("[INFO] BlueCat Logout was successful")
	return result
}

// Client logs you in to BAM and keeps your session cookie
func Client(username string, password string, endpoint string, insecure bool) (ProteusAPI, error) {
	//var response *http.Response
	cli := soap.Client{
		URL:       "https://" + endpoint + "/Services/API?wsdl",
		Namespace: Namespace,
		Pre:       setBlueCatAuthToken,
		Post:      getBlueCatAuthToken,
	}

	if insecure {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		cli.Config = client
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

// NewClient creates a client without a login
func NewClient(endpoint string, insecure bool) ProteusAPI {
	//var response *http.Response
	cli := soap.Client{
		URL:       "https://" + endpoint + "/Services/API?wsdl",
		Namespace: Namespace,
		Pre:       setBlueCatAuthToken,
		Post:      getBlueCatAuthToken,
	}

	if insecure {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		cli.Config = client
	}

	soapService := NewProteusAPI(&cli)

	return soapService
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
