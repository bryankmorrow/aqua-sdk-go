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
		detail := cli.GetRegistries()
		for _, d := range detail {
			params := make(map[string]string)
			params["include_vpatch_info"] = "true"
			params["show_negligible"] = "true"
			params["skip_count"] = "false"
			params["hide_base_image"] = "false"
			params["fix_availability"] = "true"
			params["order_by"] = "-image"
			params["registry_name"] = d.Name
			vulns := cli.GetRiskVulnerabilities(0, 0, params)
			log.Printf("Type: %s  Name: %s  Description: %s  URL: %s  Vulnerability Count: %d", d.Type, d.Name, d.Description, d.URL, vulns.Count)
		}
	}
}
