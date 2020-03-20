package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"
import (
	"encoding/json"
	"github.com/BryanKMorrow/aqua-sdk-go/types/registry"
	"github.com/parnurzeal/gorequest"
	"log"
)

// GetRegistries - retrieves all configured registry integrations
// Returns Registry struct
// Path - v1/registries
func (cli *Client) GetRegistries() registry.Registry {
	var response = registry.Registry{}
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := "/api/v1/registries"
	events, body, errs := request.Clone().Get(cli.url+apiPath).End()
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
	}
	return response
}
