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
		detail := cli.GetRiskCount()
		log.Printf("Registered Images: %d  Vulnerable Images: %d  Vulnerable Image Percent: %v\n", detail.RegisteredImages, detail.VulnerableImages, detail.VulnerableImagesPercent)
		log.Printf("Total Vulnerablities: %d  High Vulnerablities: %d  High Vulnerability Percentage of Total: %v", detail.TotalVulns, detail.HighVulns, detail.HighVulnsPercent)
	}
}
