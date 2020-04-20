package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"
import (
	"encoding/json"
	"log"

	"github.com/BryanKMorrow/aqua-sdk-go/types/gateways"
	"github.com/parnurzeal/gorequest"
)

// GetGateways retrieves the list of gateways
func (cli *Client) GetGateways() gateways.Gateways {
	var response gateways.Gateways
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := "/api/v1/servers"
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetImageAssuranceName from %s%s, %v ", cli.url, apiPath, err.Error())
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	return response
}
