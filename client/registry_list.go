package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"
import (
	"crypto/tls"
	"encoding/json"
	"github.com/BryanKMorrow/aqua-sdk-go/types/registry"
	"github.com/parnurzeal/gorequest"
	"log"
)

// GetRegistries - retrieves all configured registry integrations
// Returns Registry struct
// Path - v1/registries
func (cli *Client) GetRegistries() registry.Registry {
	var response registry.Registry
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := "/api/v1/registries"
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetAllImages from %s%s, %v ", cli.url, apiPath, err)
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	return response
}
