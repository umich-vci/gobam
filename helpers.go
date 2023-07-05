package gobam

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/fiorix/wsdl2go/soap"
)

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
	err := soapService.Login(username, password)
	if err != nil {
		return nil, fmt.Errorf("Login error: %s", err)
	}

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
