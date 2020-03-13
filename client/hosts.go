package client

import (
	"encoding/json"
	"github.com/BryanKMorrow/aqua-sdk-go/types/infrastructure"
	"github.com/parnurzeal/gorequest"
	"log"
)

// GetHosts /api/v1/hosts?hosts=
// hosts options are connected and disconnected
func (cli *Client) GetHosts(paramsString map[string]string) infrastructure.Hosts {
	var response = infrastructure.Hosts{}
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := "/api/v1/hosts"
	paramString := cli.GetStringParams(paramsString)
	events, body, errs := request.Clone().Get(cli.url + apiPath).Query(paramString).End()
	log.Printf("Calling %s%s", cli.url, apiPath)
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetHosts from %s%s, %v ", cli.url, apiPath, err.Error())
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	return response
}
