package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"

import (
	"encoding/json"
	"github.com/BryanKMorrow/aqua-sdk-go/types/enforcers"
	"github.com/parnurzeal/gorequest"
	"log"
)

// CreateEnforcerGroup - add Enforcer group
// Path - /api/v1/hostsbatch
// Payload - JSON
// Returns enforcers.GroupResponse struct
func (cli *Client) CreateEnforcerGroup(group enforcers.Group) enforcers.GroupResponse {
	var response enforcers.GroupResponse
	data, err := json.Marshal(group)
	if err != nil {
		log.Println(err)
	}
	apiPath := "/api/v1/hostsbatch"
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	resp, body, errs := request.Clone().Post(cli.url + apiPath).Send(string(data)).End()
	if errs != nil {
		log.Println(resp.StatusCode)
	}
	if resp.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func CreateEnforcerGroup from %s, %v ", cli.url, err.Error())
			//json: Unmarshal(non-pointer main.Request)
		}
	} else {
		log.Printf("Error calling func CreateEnforcerGroup from %s%s - StatusCode: %v ", cli.url, apiPath, resp.StatusCode)
	}
	return response
}
