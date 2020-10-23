package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"

import (
	"encoding/json"
	"github.com/BryanKMorrow/aqua-sdk-go/types/images"
	"github.com/parnurzeal/gorequest"
	"log"
	"strconv"
)

// GetRepositories - retrieves all configured repositories
// Returns Repository struct
// Path - api/v2/repositories
func (cli *Client) GetRepositories(page, pagesize int, paramsString map[string]string) (images.Repositories, int, int, int) {
	// set the default pagesize
	if pagesize == 0 {
		pagesize = 1000
	}
	var response images.Repositories
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := "/api/v2/repositories"
	paramString := cli.GetStringParams(paramsString)
	events, body, errs := request.Clone().Get(cli.url+apiPath).Param("page", strconv.Itoa(page)).Param("pagesize", strconv.Itoa(pagesize)).
		Query(paramString).End()
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetRespositories from %s%s, %v ", cli.url, apiPath, err)
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	remaining := cli.CalcRemaining(pagesize, page, response.Count)
	page = response.Page + 1
	remaining, next := cli.CalcNext(remaining, page)
	return response, remaining, next, response.Count
}
