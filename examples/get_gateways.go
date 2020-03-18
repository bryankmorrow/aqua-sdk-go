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
		detail := cli.GetGateways()
		for _, d := range detail {
			log.Printf("ID: %s  Logical Name: %s  Version: %s  Cluster Address: %s  Public Address: %s  GRPC Address: %s", d.ID, d.Logicalname, d.Version, d.Address, d.PublicAddress, d.GrpcAddress)
		}
	}
}
