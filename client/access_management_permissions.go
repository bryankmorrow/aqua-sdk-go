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

type PermissionSet struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	UIAccess    bool      `json:"ui_access"`
	IsSuper     bool      `json:"is_super"`
	Actions     []string  `json:"actions"`
}

type PermissionSets struct {
	Count            int             `json:"count"`
	Page             int             `json:"page"`
	Pagesize         int             `json:"pagesize"`
	Result           []PermissionSet `json:"result"`
	MoreDataAllPages int             `json:"more_data_all_pages"`
}

// GetPermissionSet - returns single Aqua permission set
func (cli *Client) GetPermissionSet(name string) (*PermissionSet, error) {
	var err error
	var response PermissionSet
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/access_management/permissions/%s", name)
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	if errs != nil {
		log.Println(events.StatusCode)
		err = fmt.Errorf("error calling %s", apiPath)
		return nil, err
	}
	if events.StatusCode == 200 {
		err = json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetPermissionSet from %s%s, %v ", cli.url, apiPath, err)
			return nil, err
		}
	}
	if response.Name == "" {
		err = fmt.Errorf("permission set not found: %s", name)
		return nil, err
	}
	return &response, err
}

// GetPermissionSets - returns all Aqua permission sets
func (cli *Client) GetPermissionSets() (*PermissionSets, error) {
	var err error
	var response PermissionSets
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/access_management/permissions")
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err = json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetPermissionSets from %s%s, %v ", cli.url, apiPath, err)
			return nil, errors.Wrap(err, "could not unmarshal permission sets response")
		}
	}
	return &response, err
}

// CreatePermissionSet - creates single Aqua permission set
func (cli *Client) CreatePermissionSet(ps PermissionSet) error {
	payload, err := json.Marshal(ps)
	if err != nil {
		return err
	}
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/access_management/permissions")
	resp, _, errs := request.Clone().Post(cli.url + apiPath).Send(string(payload)).End()
	if errs != nil {
		return errors.Wrap(err, "failed creating permission set")
	}
	if resp.StatusCode != 201 || resp.StatusCode != 204 {
		return err
	}
	return nil
}

// UpdatePermissionSet updates an existing permission set
func (cli *Client) UpdatePermissionSet(ps PermissionSet) error {
	payload, err := json.Marshal(ps)
	if err != nil {
		return err
	}
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/access_management/permissions/%s", ps.Name)
	resp, _, errs := request.Clone().Put(cli.url + apiPath).Send(string(payload)).End()
	if errs != nil {
		return errors.Wrap(err, "failed modifying permission set")
	}
	if resp.StatusCode != 201 || resp.StatusCode != 204 {
		return err
	}
	return nil
}

// DeletePermissionSet removes a permission set
func (cli *Client) DeletePermissionSet(name string) error {
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/access_management/permissions/%s", name)
	events, _, errs := request.Clone().Delete(cli.url + apiPath).End()
	if errs != nil {
		return fmt.Errorf("error while calling DELETE on /api/v2/access_management/permissions/%s: %v", name, events.StatusCode)
	}
	if events.StatusCode != 204 {
		return fmt.Errorf("failed deleting permission set, status code: %v", events.StatusCode)
	}
	return nil
}
