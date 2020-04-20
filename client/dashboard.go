package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"

import (
	"encoding/json"
	"log"

	"github.com/BryanKMorrow/aqua-sdk-go/types/dashboard"
	"github.com/parnurzeal/gorequest"
)

// GetOverview retrieves the first page dashboard statistics from v1/dashboard
// query parameters registry=&hosts=&containers_app=
// hosts doesn't do anything in this query, should call the hosts api directly
// containers_app refers to Aqua Services
// response is Overview struct
func (cli *Client) GetOverview(paramsString map[string]string) dashboard.Overview {
	var response dashboard.Overview
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := "/api/v1/dashboard"
	paramString := cli.GetStringParams(paramsString)
	events, body, errs := request.Clone().Get(cli.url + apiPath).Query(paramString).End()
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetOverview from %s%s, %v ", cli.url, apiPath, err.Error())
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	return response
}
