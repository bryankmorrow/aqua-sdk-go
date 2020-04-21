package main

import (
	"github.com/BryanKMorrow/aqua-sdk-go/client"
	"log"
	"os"
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
		detail, _, _, _ := cli.GetRepositories(0, 1000, nil)
		for _, d := range detail.Result {
			params := make(map[string]string)
			params["registry"] = d.Registry
			params["repository"] = d.Name
			images, _, _, _ := cli.GetAllImages(0, 0, params, nil)
			for _, vuln := range images.Result {
				log.Printf("Image: %s  Created: %v  Total: %d  Critical: %d  High: %d  Medium: %d  Low: %d", vuln.Name, vuln.Created, vuln.VulnsFound, vuln.CritVulns, vuln.HighVulns, vuln.MedVulns, vuln.LowVulns)
			}
		}
	}
}
