package main

import (
	"log"
	"os"

	"github.com/BryanKMorrow/aqua-sdk-go/client"
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
		params["registry"] = "Docker Hub"
		params["fix_availability"] = "true"
		vulns, _, _, _ := cli.GetAllImages(0, 0, params, nil)
		log.Println(vulns.Result)
	}
}
