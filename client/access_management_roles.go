package client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
	"log"
	"time"
)

// Role represents a local Aqua Role
type Role struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Author      string        `json:"author,omitempty"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Permission  string        `json:"permission"`
	Scopes      []string      `json:"scopes"`
	Groups      []interface{} `json:"groups,omitempty"`
	Users       []string      `json:"users,omitempty"`
}

// Roles represents the API return for multiple roles
type Roles struct {
	Count            int    `json:"count"`
	Page             int    `json:"page"`
	Pagesize         int    `json:"pagesize"`
	Result           []Role `json:"result"`
	MoreDataAllPages int    `json:"more_data_all_pages"`
}

// GetRole - returns single Aqua Role
func (cli *Client) GetRole(name string) (*Role, error) {
	var err error
	var response Role
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/access_management/roles/%s", name)
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	if errs != nil {
		log.Println(events.StatusCode)
		err = fmt.Errorf("error calling %s", apiPath)
		return nil, err
	}
	if events.StatusCode == 200 {
		err = json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetRole from %s%s, %v ", cli.url, apiPath, err)
			return nil, err
		}
	}
	if response.Name == "" {
		err = fmt.Errorf("role not found: %s", name)
		return nil, err
	}
	return &response, err
}

// GetRoles - returns all Aqua Roles
func (cli *Client) GetRoles() ([]Role, error) {
	var err error
	var response []Role
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/access_management/roles")
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err = json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetRoles from %s%s, %v ", cli.url, apiPath, err)
			return nil, errors.Wrap(err, "could not unmarshal Roles response")
		}
	}
	return response, err
}

// CreateRole - creates single Aqua Role
func (cli *Client) CreateRole(role Role) error {
	payload, err := json.Marshal(role)
	if err != nil {
		return err
	}
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/access_management/roles")
	resp, _, errs := request.Clone().Post(cli.url + apiPath).Send(string(payload)).End()
	if errs != nil {
		return errors.Wrap(err, "failed creating role")
	}
	if resp.StatusCode != 201 || resp.StatusCode != 204 {
		return err
	}
	return nil
}

// UpdateRole updates an existing Role
func (cli *Client) UpdateRole(role Role) error {
	payload, err := json.Marshal(role)
	if err != nil {
		return err
	}
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/access_management/roles/%s", role.Name)
	resp, _, errs := request.Clone().Put(cli.url + apiPath).Send(string(payload)).End()
	if errs != nil {
		return errors.Wrap(err, "failed modifying role")
	}
	if resp.StatusCode != 201 || resp.StatusCode != 204 {
		return err
	}
	return nil
}

// DeleteRole removes a Role
func (cli *Client) DeleteRole(name string) error {
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/access_management/roles/%s", name)
	events, _, errs := request.Clone().Delete(cli.url + apiPath).End()
	if errs != nil {
		return fmt.Errorf("error while calling DELETE on /api/v2/access_management/roles/%s: %v", name, events.StatusCode)
	}
	if events.StatusCode != 204 {
		return fmt.Errorf("failed deleting Role, status code: %v", events.StatusCode)
	}
	return nil
}
