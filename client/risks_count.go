package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"
import (
	"encoding/json"
	"github.com/BryanKMorrow/aqua-sdk-go/types/risks"
	"github.com/parnurzeal/gorequest"
	"log"
)

// GetRiskCount - retrieves risk counts
// Returns risks.Counts struct
// Path - v2/risks
func (cli *Client) GetRiskCount() risks.Counts {
	var response risks.Counts
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := "/api/v2/risks"
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetRiskCount from %s%s, %v ", cli.url, apiPath, err)
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	return response
}
