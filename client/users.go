package client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
	"log"
)

// User represents a local Aqua user
type User struct {
	ID              string   `json:"id"` // Username
	Password        string   `json:"password,omitempty"`
	PasswordConfirm string   `json:"passwordConfirm,omitempty"`
	Roles           []string `json:"roles,omitempty"`
	Name            string   `json:"name,omitempty"` // Display Name
	Email           string   `json:"email,omitempty"`
	FirstTime       bool     `json:"first_time,omitempty"`
}

// GetUser - returns single Aqua user
// Params: name: The name of user
func (cli *Client) GetUser(name string) (User, error) {
	var err error
	var response User
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
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
func (cli *Client) GetUsers() ([]User, error) {
	var err error
	var response []User
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
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
func (cli *Client) CreateUser(user User) error {
	payload, err := json.Marshal(user)
	if err != nil {
		return err
	}
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
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

// UpdateUser updates an existing user
func (cli *Client) UpdateUser(name string) (*User, error) {
	res := &User{}

	return res, nil
}

// DeleteUser
func (cli *Client) DeleteUser(name string) error {

	return nil
}
