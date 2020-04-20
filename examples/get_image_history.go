package main

import (
	"encoding/json"
	"github.com/BryanKMorrow/aqua-sdk-go/client"
	"log"
	"os"
	"sort"
	"time"
)

type ImageHistory struct {
	Name     string            `json:"name"`
	Registry string            `json:"registry"`
	Tag      []ImageHistoryTag `json:"image_history_tag"`
}

type ImageHistoryTag struct {
	Tag        string    `json:"tag"`
	Created    time.Time `json:"created"`
	VulnsFound int       `json:"vulns_found"`
	CritVulns  int       `json:"crit_vulns"`
	HighVulns  int       `json:"high_vulns"`
	MedVulns   int       `json:"med_vulns"`
	LowVulns   int       `json:"low_vulns"`
	NegVulns   int       `json:"neg_vulns"`
}

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
		params["registry"] = "Docker Hub"
		params["repository"] = "httpd"
		params["fix_availability"] = "true"
		img := ImageHistory{
			Name:     params["repository"],
			Registry: params["registry"],
		}
		var iht []ImageHistoryTag
		images, _, _, _ := cli.GetAllImages(0, 0, params, nil)
		for _, image := range images.Result {
			//log.Printf("Image: %s  Created: %v  Total: %d  Critical: %d  High: %d  Medium: %d  Low: %d", vuln.Name, vuln.Created, vuln.VulnsFound, vuln.CritVulns, vuln.HighVulns, vuln.MedVulns, vuln.LowVulns)
			i, err := cli.GetImage(image.Registry, image.Repository, image.Tag)
			if err != nil {
				log.Println(err)
			} else {
				tag := ImageHistoryTag{
					Tag:        i.Tag,
					Created:    i.Metadata.Created,
					VulnsFound: i.VulnsFound,
					CritVulns:  i.CritVulns,
					HighVulns:  i.HighVulns,
					MedVulns:   i.MedVulns,
					LowVulns:   i.LowVulns,
					NegVulns:   i.NegVulns,
				}
				iht = append(iht, tag)
				//log.Printf("Image %s created on: %v", i.Name, i.Metadata.Created)
				//log.Printf("Total: %d  Critical: %d  High: %d  Medium: %d  Low: %d", i.VulnsFound, i.CritVulns, i.HighVulns, i.MedVulns, i.LowVulns)
			}
		}
		sort.Slice(iht, func(i, j int) bool {
			return iht[i].Created.Before(iht[j].Created)
		})
		img.Tag = iht
		data, _ := json.Marshal(img)
		log.Println(string(data))
	}
}
