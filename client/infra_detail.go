package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/BryanKMorrow/aqua-sdk-go/types/infrastructure"
	"github.com/parnurzeal/gorequest"
	"log"
)

// GetDetailName - gets a single infrastructure asset by name
// Path v2/infrastructure/<type>/<name> (type can be cluster or node
// Return - infrastructure.Detail response struct
func (cli *Client) GetDetailName(infraType, name string) infrastructure.Detail {
	var response infrastructure.Detail
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/infrastructure/%s/%s", infraType, name)
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetDetailName from %s%s, %v ", cli.url, apiPath, err)
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	return response
}

// GetDetailID - gets a single infrastructure asset by ID
// Path v2/infrastructure/<id>
// Return - infrastructure.Detail response struct
// NOTE - may only work for Nodes, not Clusters??
func (cli *Client) GetDetailID(ID string) infrastructure.Detail {
	var response = infrastructure.Detail{}
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/infrastructure/%s", ID)
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetDetailID from %s%s, %v ", cli.url, apiPath, err)
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	return response
}
