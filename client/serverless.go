package client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
	"log"
)

// ServerlessProject allows the configuration of a function/serverless integration
type ServerlessProject struct {
	ID                 int      `json:"id,omitempty"`
	Name               string   `json:"name"`
	Description        string   `json:"description,omitempty"`
	Username           string   `json:"username"`
	Password           string   `json:"password,omitempty"`
	Region             string   `json:"region"`
	ComputeProvider    int      `json:"compute_provider"` // AWS = 1, Azure = 3
	Author             string   `json:"author"`
	Update             int      `json:"update,omitempty"`
	AutoPull           bool     `json:"auto_pull,omitempty"`
	AutoPullTime       string   `json:"auto_pull_time,omitempty"`
	AutoPullInProgress bool     `json:"auto_pull_in_progress,omitempty"`
	SqsURL             string   `json:"sqs_url,omitempty"`
	IncludeTags        []string `json:"include_tags,omitempty"`
	ExcludeTags        []string `json:"exclude_tags,omitempty"`
}

// GetServerlessProject - returns single Aqua Serverless Function Project
func (cli *Client) GetServerlessProject(name string) (*ServerlessProject, error) {
	var err error
	var response ServerlessProject
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/serverless/projects/%s", name)
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	if errs != nil {
		log.Println(events.StatusCode)
		err = fmt.Errorf("error calling %s", apiPath)
		return nil, err
	}
	if events.StatusCode == 200 {
		err = json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetServerlessProject from %s%s, %v ", cli.url, apiPath, err)
			return nil, err
		}
	}
	if response.Name == "" {
		err = fmt.Errorf("serverless project not found: %s", name)
		return nil, err
	}
	return &response, err
}

// GetServerlessProjects is the return of all serverless projects
func (cli *Client) GetServerlessProjects() ([]ServerlessProject, error) {
	var err error
	var response []ServerlessProject
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/serverless/projects")
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err = json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetServerlessProjects from %s%s, %v ", cli.url, apiPath, err)
			return nil, errors.Wrap(err, "could not unmarshal Serverless Projects response")
		}
	}
	return response, err
}

// CreateServerlessProject - creates single Aqua Serverless Function Project
func (cli *Client) CreateServerlessProject(proj ServerlessProject) error {
	payload, err := json.Marshal(proj)
	if err != nil {
		return err
	}
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/serverless/projects")
	resp, _, errs := request.Clone().Post(cli.url + apiPath).Send(string(payload)).End()
	if errs != nil {
		return errors.Wrap(err, "failed creating serverless project")
	}
	if resp.StatusCode != 201 || resp.StatusCode != 204 {
		return err
	}
	return nil
}

// UpdateServerlessProject updates an existing Serverless Function Project
func (cli *Client) UpdateServerlessProject(proj ServerlessProject) error {
	payload, err := json.Marshal(proj)
	if err != nil {
		return err
	}
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/serverless/projects/%s", proj.Name)
	resp, _, errs := request.Clone().Put(cli.url + apiPath).Send(string(payload)).End()
	if errs != nil {
		return errors.Wrap(err, "failed modifying serverless project")
	}
	if resp.StatusCode != 201 || resp.StatusCode != 204 {
		return err
	}
	return nil
}

// DeleteServerlessProject removes a Serverless Project
func (cli *Client) DeleteServerlessProject(name string) error {
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/serverless/project/%s", name)
	events, _, errs := request.Clone().Delete(cli.url + apiPath).End()
	if errs != nil {
		return fmt.Errorf("error while calling DELETE on /api/v2/serverless/project/%s: %v", name, events.StatusCode)
	}
	if events.StatusCode != 204 {
		return fmt.Errorf("failed deleting serverless project, status code: %v", events.StatusCode)
	}
	return nil
}
