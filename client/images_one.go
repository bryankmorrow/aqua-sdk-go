package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/BryanKMorrow/aqua-sdk-go/types/images"
	"github.com/parnurzeal/gorequest"
)

// GetImage -  retrieves a single image based on registry, repo and tag
// Accepts the registry, repo and tag
// Returns response struct
// v2/images/registry/repo/tag
func (cli *Client) GetImage(registry, repo, tag string) (images.SingleResponse, error) {
	var response = images.SingleResponse{}
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/images/%s/%s/%s", registry, repo, tag)
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	log.Printf("Calling %s%s", cli.url, apiPath)
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetAllImages from %s%s, %v ", cli.url, apiPath, err.Error())
			//json: Unmarshal(non-pointer main.Request)
		}
	} else {
		return response, errors.New("image not found")
	}
	return response, nil
}
