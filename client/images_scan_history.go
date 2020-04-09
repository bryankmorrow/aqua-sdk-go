package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/BryanKMorrow/aqua-sdk-go/types/images"
	"github.com/parnurzeal/gorequest"
)

// GetScanHistory -  retrieves a single image scan history based on registry, repo and tag
// Path parameters of {registry}, {repo} and {tag}
// Returns response struct
// v2/images/{registry}/{repo}/{tag}/scan_history
func (cli *Client) GetScanHistory(registry, repo, tag string) (images.ScanHistory, error) {
	var response = images.ScanHistory{}
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/images/%s/%s/%s/scan_history", registry, repo, tag)
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	log.Printf("Calling %s%s", cli.url, apiPath)
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetScanHistory from %s%s, %v ", cli.url, apiPath, err.Error())
			//json: Unmarshal(non-pointer main.Request)
		}
	} else {
		return response, errors.New("image not found")
	}
	return response, nil
}
