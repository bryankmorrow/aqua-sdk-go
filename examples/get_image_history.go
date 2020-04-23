package main

import (
	"encoding/json"
	"github.com/BryanKMorrow/aqua-sdk-go/client"
	"github.com/BryanKMorrow/aqua-sdk-go/types/images"
	"log"
	"os"
	"sort"
)

func main() {
	// Get the Aqua CSP connection parameters from Environment Variables
	url := os.Getenv("AQUA_URL")
	user := os.Getenv("AQUA_USER")
	password := os.Getenv("AQUA_PASSWORD")

	// Create the client and get the JWT token for API call authorization
	cli := client.NewClient(url, user, password)
	connected := cli.GetAuthToken()
	if !connected {
		log.Fatalln("Failed to retrieve JWT Authorization Token")
	} else {

		params := make(map[string]string)
		params["registry"] = "registry"
		params["repository"] = "repository"
		params["fix_availability"] = "true"
		img := images.History{
			Name:     params["repository"],
			Registry: params["registry"],
		}
		var tl []images.Tag
		il, _, _, _ := cli.GetAllImages(0, 0, params, nil)
		for _, image := range il.Result {
			//log.Printf("Image: %s  Created: %v  Total: %d  Critical: %d  High: %d  Medium: %d  Low: %d", vuln.Name, vuln.Created, vuln.VulnsFound, vuln.CritVulns, vuln.HighVulns, vuln.MedVulns, vuln.LowVulns)

			i, err := cli.GetImage(image.Registry, image.Repository, image.Tag)
			if err == nil {
				tag := images.Tag{
					Tag:        i.Tag,
					Created:    i.Metadata.Created,
					VulnsFound: i.VulnsFound,
					CritVulns:  i.CritVulns,
					HighVulns:  i.HighVulns,
					MedVulns:   i.MedVulns,
					LowVulns:   i.LowVulns,
					NegVulns:   i.NegVulns,
				}
				tl = append(tl, tag)
				//log.Printf("Image %s created on: %v", i.Name, i.Metadata.Created)
				//log.Printf("Total: %d  Critical: %d  High: %d  Medium: %d  Low: %d", i.VulnsFound, i.CritVulns, i.HighVulns, i.MedVulns, i.LowVulns)
			}
		}
		sort.Slice(tl, func(i, j int) bool {
			return tl[i].Created.Before(tl[j].Created)
		})
		img.Tags = tl
		data, _ := json.Marshal(img)
		log.Println(string(data))
	}
}
