package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"
import (
	"encoding/json"
	"github.com/BryanKMorrow/aqua-sdk-go/types/assurance"
	"github.com/parnurzeal/gorequest"
	"log"
)

// GetAssurancePolicies - gets all Assurance Policy types
// Params: identifiers_only bool (true|false), order_by string (name)
// Returns: Struct from types/assurance/policies
func (cli *Client) GetAssurancePolicies(paramsString map[string]string) assurance.Policies {
	var response assurance.Policies
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := "/api/v2/assurance_policy"
	paramString := cli.GetStringParams(paramsString)
	log.Printf(paramString)
	events, body, errs := request.Clone().Get(cli.url + apiPath).Query(paramString).End()
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetAssurancePolicies from %s%s, %v ", cli.url, apiPath, err)
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	return response
}
