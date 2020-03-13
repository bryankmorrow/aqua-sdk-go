package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"
import (
	"encoding/json"
	"github.com/BryanKMorrow/aqua-sdk-go/types/assurance"
	"github.com/parnurzeal/gorequest"
	"log"
)

// GetAssuranceScripts - retrieves user created assurance scripts
// Path - /api/v2/image_assurance/user_scripts
// Query Params - search, order_by (name|num_users), engine (ab|sh|yaml)
// Returns assurance.Scripts struct
func (cli *Client) GetAssuranceScripts(paramsString  map[string]string) assurance.Scripts {
	var response = assurance.Scripts{}
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := "/api/v2/image_assurance/user_scripts"
	paramString := cli.GetStringParams(paramsString)
	events, body, errs := request.Clone().Get(cli.url + apiPath).Query(paramString).End()
	log.Printf("Calling %s%s", cli.url, apiPath)
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetAssuranceScripts from %s%s, %v ", cli.url, apiPath, err.Error())
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	return response
}
