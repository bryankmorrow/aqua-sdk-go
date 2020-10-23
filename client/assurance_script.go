package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"
import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/BryanKMorrow/aqua-sdk-go/types/assurance"
	"github.com/parnurzeal/gorequest"
)

// GetAssuranceScriptID - retrieves user created assurance script by ID
// Params: ID: string representation of the script ID (0,1,2 etc)
// Path - /api/v2/image_assurance/user_scripts/<ID>
// Returns: Struct from types/assurance/script
func (cli *Client) GetAssuranceScriptID(ID string) assurance.Script {
	var response assurance.Script
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/image_assurance/user_scripts/%s", ID)
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetAssuranceScriptID from %s%s, %v ", cli.url, apiPath, err)
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	return response
}
