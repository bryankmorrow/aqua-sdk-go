package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/BryanKMorrow/aqua-sdk-go/types/images"
	"github.com/parnurzeal/gorequest"
)

// GetImage -  retrieves a single image based on registry, repo and tag
// Accepts the registry, repo and tag
// Returns Image struct
// v2/images/registry/repo/tag
func (cli *Client) GetImage(registry, repo, tag string) (images.Image, error) {
	var response images.Image
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/images/%s/%s/%s", registry, repo, tag)
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetImage from %s%s, %v ", cli.url, apiPath, err)
			if terr, ok := err.(*time.ParseError); ok {
				fmt.Println("Error when trying to read 'metadata.created' value", terr)
				log.Println("Setting the metadata.created to current timestamp to bypass this error")
				response.Metadata.Created = time.Now()
			}
		}
	} else {
		return response, errors.New("image not found")
	}
	return response, nil
}
