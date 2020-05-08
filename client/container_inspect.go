package client

import (
	"encoding/json"
	"fmt"
	"github.com/BryanKMorrow/aqua-sdk-go/types/containers"
	"github.com/parnurzeal/gorequest"
	"log"
)

// InspectContainer - retrieves metadata from running container
// Accepts container id and host id
// Returns response struct
// v1/containers/{id}/{host_id}/inspect
func (cli *Client) InspectContainer(ID, hostID string) containers.Inspect {
	var response containers.Inspect
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v1/containers/%s/%s/inspect", ID, hostID)
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func InspectContainer from %s%s, %v ", cli.url, apiPath, err)
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	return response
}
