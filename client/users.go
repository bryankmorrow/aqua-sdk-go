package client

import (
	"encoding/json"
	"fmt"
	"github.com/BryanKMorrow/aqua-sdk-go/types/users"
	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
	"log"
)

// GetUser - returns single Aqua user
// Params: name: The name of user
func (cli *Client) GetUser(name string) (users.User, error) {
	var err error
	var response users.User
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v1/users/%s", name)
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err = json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetUser from %s%s, %v ", cli.url, apiPath, err)
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	if response.Name == "" {
		err = fmt.Errorf("user not found: %s", name)
	}
	return response, err
}

// GetUsers - returns all Aqua users
func (cli *Client) GetUsers() ([]users.User, error) {
	var err error
	var response []users.User
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v1/users")
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err = json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetUser from %s%s, %v ", cli.url, apiPath, err)
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	return response, err
}

// CreateUser - creates single Aqua user
// Params: name: The name of the Image Assurance Policy
// Returns: The struct from types/users/user
func (cli *Client) CreateUser(user users.User) error {
	payload, err := json.Marshal(user)
	if err != nil {
		return err
	}
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v1/users")
	resp, _, errs := request.Clone().Post(cli.url + apiPath).Send(string(payload)).End()
	if errs != nil {
		return errors.Wrap(err, "failed creating user")
	}
	if resp.StatusCode != 201 || resp.StatusCode != 204 {
		return err
	}
	return nil
}
