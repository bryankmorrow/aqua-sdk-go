package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/BryanKMorrow/aqua-sdk-go/types/images"
	"github.com/parnurzeal/gorequest"
)

// GetLayers - retrieves all layers from a particular image
// Accepts the registry, repo and tag strings as well as the page number, pagesize and params map
// Returns response struct, remaining count and next page
// v2/images
func (cli *Client) GetLayers(registry, repo, tag string, page, pagesize int, paramsString map[string]string, paramsBool map[string]bool) (images.Layers, int, int, int) {
	// set the default pagesize
	if pagesize == 0 {
		pagesize = 1000
	}
	var response = images.Layers{}
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/images/%s/%s/%s/history_layers", registry, repo, tag)
	paramString := cli.GetStringParams(paramsString)
	paramBool := cli.GetBoolParams(paramsBool)
	events, body, errs := request.Clone().Get(cli.url+apiPath).Param("page", strconv.Itoa(page)).Param("pagesize", strconv.Itoa(pagesize)).
		Query(paramString).Query(paramBool).End()
	log.Printf("Calling %s%s", cli.url, apiPath)
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetLayers from %s%s, %v ", cli.url, apiPath, err.Error())
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	// remove duplicate layers
	newLayers := removeDuplicates(response)
	// create new Layer{} for response
	var newResponse = images.Layers{
		Count:    response.Count,
		Page:     response.Page,
		Pagesize: response.Pagesize,
		Result:   newLayers,
	}

	remaining := cli.CalcRemaining(pagesize, page, response.Count)
	page = response.Page + 1
	remaining, next := cli.CalcNext(remaining, page)
	return newResponse, remaining, next, response.Count
}

func removeDuplicates(response images.Layers) images.LayerResult {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := images.LayerResult{}

	for _, layer := range response.Result {
		if encountered[layer.ID] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[layer.ID] = true
			// Append to result slice.
			result = append(result, layer)
		}
	}
	// Return the new slice.
	return result
}
