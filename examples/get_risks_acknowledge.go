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
		params := make(map[string]string)
		params["text_search"] = "CVE-2020"
		detail, remaining, next, total := cli.GetRisksAcknowledge(params)
		log.Printf("Remaining: %d  Next Page: %d  Total: %d\n", remaining, next, total)
		for _, d := range detail.Result {
			log.Printf("Issue Type: %s  Issue Name: %s  Resource Type: %s  Resource Name: %s  Resource Version: %s\n", d.IssueType, d.IssueName, d.ResourceType, d.ResourceName, d.ResourceVersion)
		}
	}
}
