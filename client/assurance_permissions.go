package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"
import (
	"encoding/json"
	"github.com/BryanKMorrow/aqua-sdk-go/types/assurance"
	"github.com/parnurzeal/gorequest"
	"log"
)

// GetImagePermissions returns the globally whitelisted and blacklisted images
// Response is PermissionList struct
func (cli *Client) GetImagePermissions() assurance.PermissionList {
	var response = assurance.PermissionList{}
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := "/api/v2/image_assurance/image_permissions"
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	log.Printf("Calling %s%s", cli.url, apiPath)
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetImagePermissions from %s%s, %v ", cli.url, apiPath, err.Error())
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	return response
}
