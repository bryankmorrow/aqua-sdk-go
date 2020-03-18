package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"

import (
	"encoding/json"
	"github.com/BryanKMorrow/aqua-sdk-go/types/enforcers"
	"github.com/parnurzeal/gorequest"
	"log"
)

// GetEnforcerGroups - retrieve list of Enforcer Groups
// Path - GET /api/v1/hostsbatch
// Returns a slice of enforcers.GroupResponse
func (cli *Client) GetEnforcerGroups() []enforcers.GroupResponse {
	var response []enforcers.GroupResponse
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := "/api/v1/hostsbatch"
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	log.Printf("Calling %s%s", cli.url, apiPath)
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetEnforcerGroups from %s%s, %v ", cli.url, apiPath, err.Error())
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	return response
}
