package client

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
	"log"
	"time"
)

// ASVariable is the generic variable payload for Application Scoping Categorizes
type ASVariable struct {
	Attribute string `json:"attribute"`
	Value     string `json:"value"`
}

// ASCategory is the generic struct for Application Scope Categories
type ASCategory struct {
	Expression string       `json:"expression"`
	Variables  []ASVariable `json:"variables"`
}

// ASArtifacts is an Application Scope category
type ASArtifacts struct {
	Image    ASCategory `json:"image"`
	Function ASCategory `json:"function"`
	Cf       ASCategory `json:"cf"`
}

// ASWorkloads is an Application Scope category
type ASWorkloads struct {
	Kubernetes ASCategory `json:"kubernetes"`
	Os         ASCategory `json:"os"`
	Cf         ASCategory `json:"cf"`
}

// Infrastructure is an Application Scope category
type ASInfrastructure struct {
	Kubernetes ASCategory `json:"kubernetes"`
	Os         ASCategory `json:"os"`
}

// ApplicationScope is the payload for creating an application scope
// URL: /api/v2/access_management/scopes
type ApplicationScope struct {
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	OwnerEmail  string    `json:"owner_email,omitempty"`
	Author      string    `json:"author,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Categories  struct {
		Artifacts      ASArtifacts      `json:"artifacts"`
		Workloads      ASWorkloads      `json:"workloads"`
		Infrastructure ASInfrastructure `json:"infrastructure"`
	} `json:"categories"`
}

// ReadApplicationScope retrieves an application scope from the Aqua Enterprise API
func (cli *Client) ReadApplicationScope(name string) (*ApplicationScope, error) {
	as := &ApplicationScope{}
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/access_management/scopes/%s", name)
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	if errs != nil {
		log.Println("[DEBUG]  failed to get application scopes: ", events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), as)
		if err != nil {
			return nil, errors.Wrap(err, "could not unmarshal application scope response")
		}
	}
	return as, nil
}

// CreateApplicationScope creates a new application scope in Aqua Enterprise API
func (cli *Client) CreateApplicationScope(scope *ApplicationScope) error {
	payload, err := json.Marshal(scope)
	if err != nil {
		return err
	}
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/access_management/scopes")
	resp, _, errs := request.Clone().Post(cli.url + apiPath).Send(string(payload)).End()
	if errs != nil {
		return errors.Wrap(err, "failed creating application scope")
	}
	if resp.StatusCode != 201 || resp.StatusCode != 204 {
		return err
	}
	return nil
}

// UpdateApplicationScope updates an existing application scope in Aqua Enterprise API
func (cli *Client) UpdateApplicationScope(name string, req ApplicationScope) (*ApplicationScope, error) {
	res := &ApplicationScope{}

	return res, nil
}

// DeleteApplicationScope
func (cli *Client) DeleteApplicationScope(name string) error {

	return nil
}
