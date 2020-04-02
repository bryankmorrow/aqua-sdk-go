package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"

import (
	"encoding/json"
	"github.com/BryanKMorrow/aqua-sdk-go/types/risks"
	"github.com/parnurzeal/gorequest"
	"log"
	"strconv"
)

// GetRisksAcknowledge retrieves a list of vulnerabilities whose risk has been accepted
// Param: page: int - Page to retrieve
// Param: page_size: int - number of elements to retrieve
// Param: fix_availability: bool - retrieve only elements which do or do not have a fix version available
// Param: text_search: string - search and retrieve elements that match search string
// Param: order_by: string - order based on specific fields (-field to reverse the order)
func (cli *Client) GetRisksAcknowledge(paramsString map[string]string) (risks.Acknowledge, int, int, int) {
	// convert paramsString["page"] to int
	page, _ := strconv.Atoi(paramsString["page"])
	// set the default pagesize
	pagesize, _ := strconv.Atoi(paramsString["page_size"])
	if pagesize == 0 {
		paramsString["page_size"] = "1000"
		pagesize = 1000
	}
	var response = risks.Acknowledge{}
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := "/api/v2/risks/acknowledge"
	paramString := cli.GetStringParams(paramsString)
	events, body, errs := request.Clone().Get(cli.url + apiPath).Query(paramString).End()
	//log.Printf("Calling %s%s", cli.url, apiPath)
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetRisksAcknowledge from %s%s, %v ", cli.url, apiPath, err.Error())
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	remaining := cli.CalcRemaining(pagesize, page, response.Count)
	page = response.Page + 1
	remaining, next := cli.CalcNext(remaining, page)
	return response, remaining, next, response.Count
}
