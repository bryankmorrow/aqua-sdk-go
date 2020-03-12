package client

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/BryanKMorrow/aqua-sdk-go/types/images"

	"github.com/parnurzeal/gorequest"
)

// GetAllImages - retrieves all registered Images with params and pagination
// Accepts the CSP struct, page number, pagesize and params map
// Returns response struct, remaining count and next page
// v2/images
func (cli *Client) GetAllImages(page, pagesize int, params map[string]string) (images.AllResponse, int, int, int) {
	// set the default pagesize
	if pagesize == 0 {
		pagesize = 1000
	}
	var response = images.AllResponse{}
	var paramString string
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := "/api/v2/images"
	// iterate through the map of params
	i := 1
	for k, v := range params {
		if i == 1 {
			paramString = k + "=" + v
		} else {
			paramString = paramString + "&" + k + "=" + v
		}
		i++
	}
	events, body, errs := request.Clone().Get(cli.url+apiPath).Param("page", strconv.Itoa(page)).Param("pagesize", strconv.Itoa(pagesize)).Query(paramString).End()
	log.Printf("Calling %s%s", cli.url, apiPath)
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetAllImages from %s%s, %v ", cli.url, apiPath, err.Error())
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	remaining := cli.CalcRemaining(pagesize, page, response.Count)
	page = response.Page + 1
	remaining, next := cli.CalcNext(remaining, page)
	return response, remaining, next, response.Count
}
