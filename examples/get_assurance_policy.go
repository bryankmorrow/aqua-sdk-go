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
		// Get assurance policy identifiers only, not the detailed data.
		params := make(map[string]string)
		params["identifiers_only"] = "true"
		// Get a list of all assurance policy types (Image, Host, Function, CF)
		detail := cli.GetAssurancePolicies(params)
		for _, d := range detail.Result {
			log.Printf("Type: %s   Name: %s  Description: %s  Last Updated: %v", d.AssuranceType, d.Name, d.Description, d.Lastupdate)
		}
	}
}
