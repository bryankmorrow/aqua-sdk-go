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
		vulns, _, _, _ := cli.GetVulnerabilities("Host Images", "registry.redhat.io/openshift4/ose-ansible-service-broker-operator", "@sha256:1cc1eac9cc94bed887ee70d4649f01cdcbc7bcef2eccae3ca0c483d3e622f39d", 0, 1000, nil, nil)
		log.Println(vulns)
	}
}
