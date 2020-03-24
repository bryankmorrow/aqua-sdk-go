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
		// Get a single image
		detail, err := cli.GetImage("registry", "repository", "tag")
		if err == nil {
			log.Printf("Image: %s  Non-Compliant: %v  Total Vulns: %d  Critical Vulns: %d  High Vulns: %d", detail.Name, detail.Disallowed, detail.VulnsFound, detail.CritVulns, detail.HighVulns)
		} else {
			log.Printf(err.Error())
		}
	}
}
