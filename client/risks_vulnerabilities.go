package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"

import (
	"encoding/json"
	"github.com/BryanKMorrow/aqua-sdk-go/types/risks"
	"github.com/parnurzeal/gorequest"
	"log"
	"strconv"
)

// GetRiskVulnerabilities - retrieves all at risk vulnerabilities
// Returns Repository struct
// Path - api/v2/repositories
func (cli *Client) GetRiskVulnerabilities(page, pagesize int, paramsString map[string]string) risks.Vulnerabilities {
	// set the default pagesize
	if pagesize == 0 {
		pagesize = 1000
	}
	var response risks.Vulnerabilities
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := "/api/v2/risks/vulnerabilities"
	paramString := cli.GetStringParams(paramsString)
	events, body, errs := request.Clone().Get(cli.url+apiPath).Param("page", strconv.Itoa(page)).Param("pagesize", strconv.Itoa(pagesize)).
		Query(paramString).End()
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetRiskVulnerabilities from %s%s, %v ", cli.url, apiPath, err.Error())
			//json: Unmarshal(non-pointer main.Request)
		}
	}

	return response
}
