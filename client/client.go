/*
This is a Golang client for the Aqua Cloud Native Security REST API.

Usage

*/

package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"

import (
	"encoding/json"
	"log"

	"github.com/parnurzeal/gorequest"
)

// Client is the REST API client that communicates with Aqua cli
type Client struct {
	// url is the base path to Aqua API
	url string
	// user that is used for the Aqua API connection
	user string
	// password is the Aqua API user's password
	password string
	// token is the JWT token received after basic authentication
	token string
	// name is the logical identifier given to the connection
	name string
}

// NewClient - initialize and return the Client
func NewClient(url, user, password string) *Client {
	c := &Client{
		url:      url,
		user:     user,
		password: password,
	}
	return c
}

// GetAuthToken - Connect to Aqua and return a JWT bearerToken (string)
// Return: bool - successfully connected?
func (cli *Client) GetAuthToken() bool {
	var connected bool
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, body, err := request.Post(cli.url+"/api/v1/login").Param("abilities", "1").
		Send(`{"id":"` + cli.user + `", "password":"` + cli.password + `"}`).End()

	if err != nil {
		log.Printf("Failed connecting to Aqua cli: %s \n  Status Code: %d", cli.url, resp.StatusCode)
		connected = false
		return connected
	}

	if resp.StatusCode == 200 {
		var raw map[string]interface{}
		_ = json.Unmarshal([]byte(body), &raw)
		cli.token = raw["token"].(string)
		connected = true
	} else {
		log.Printf("Failed with status: %s", resp.Status)
		connected = false
	}
	return connected
}
