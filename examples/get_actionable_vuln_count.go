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
		params["include_vpatch_info"] = "true"
		params["show_negligible"] = "true"
		params["skip_count"] = "false"
		params["hide_base_image"] = "false"
		params["fix_availability"] = "true"
		params["order_by"] = "-image"
		detail, _, _, _ := cli.GetRiskVulnerabilities(1, 50, params)
		log.Println(detail.Count)
	}
}
