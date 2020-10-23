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
		// Return only engine type sh
		params := make(map[string]string)
		params["engine"] = "sh"
		// Get custom compliance script information
		detail := cli.GetAssuranceScripts(params)
		// Whitelist
		for _, d := range detail {
			log.Printf("Name: %s   Description: %s  Path: %s", d.Name, d.Description, d.Path)
		}
	}
}
