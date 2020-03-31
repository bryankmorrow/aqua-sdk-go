package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"

import (
	"encoding/json"
	"fmt"
	"github.com/BryanKMorrow/aqua-sdk-go/types/assurance"
	"github.com/parnurzeal/gorequest"
	"log"
)

// GetImageAssurance - This returns the Image Assurance Policy by name
// Params: name: The name of the Image Assurance Policy
// Returns: The struct from types/assurance/image
func (cli *Client) GetImageAssuranceName(name string) (assurance.Image, error) {
	var err error
	var response = assurance.Image{}
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/image_assurance/%s", name)
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	log.Printf("Calling %s%s", cli.url, apiPath)
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err = json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetImageAssuranceName from %s%s, %v ", cli.url, apiPath, err.Error())
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	if response.Name == "" {
		err = fmt.Errorf("image assurance policy not found: %s", name)
	}
	return response, err
	}

