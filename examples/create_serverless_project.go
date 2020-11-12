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
		reg := client.ServerlessProject{
			Name:            "terraform-demo",
			Description:     "description",
			ComputeProvider: 1,
			Region:          "us-east-1",
			Username:        "APIKEY",
			Password:        "SECRETKEY",
			AutoPull:        false,
			AutoPullTime:    "00:00",
		}
		err := cli.CreateServerlessProject(reg)
		if err != nil {
			log.Println(err)
		}
	}
}
