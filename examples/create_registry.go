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
		reg := client.Registry{
			Name:        "terraform",
			Type:        "HUB",
			URL:         "https://docker.io",
			Author:      "Terraform",
			Description: "terraform creation ",
		}
		err := cli.CreateRegistry(reg)
		if err != nil {
			log.Println(err)
		}
	}
}
