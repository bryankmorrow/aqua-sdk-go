package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"

	"github.com/BryanKMorrow/aqua-sdk-go/types/dashboard"
	"github.com/parnurzeal/gorequest"
)

// GetTrends retrieves the first page dashboard trends from v1/dashboard/<trend>/trends
// path parameters are containers,images,vulnerabilities
// response is Trends struct
func (cli *Client) GetTrends(trend string) dashboard.Trends {
	var response dashboard.Trends
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v1/dashboard/%s/trends", trend)
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetTrends from %s%s, %v ", cli.url, apiPath, err)
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	return response
}
