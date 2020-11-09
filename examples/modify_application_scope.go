package main

import (
	"github.com/BryanKMorrow/aqua-sdk-go/client"
	"log"
	"os"
)

type Category struct {
	Artifacts      client.ASArtifacts      `json:"artifacts"`
	Workloads      client.ASWorkloads      `json:"workloads"`
	Infrastructure client.ASInfrastructure `json:"infrastructure"`
}

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
		artifacts := client.ASArtifacts{
			Image:    client.ASCategory{},
			Function: client.ASCategory{},
			Cf:       client.ASCategory{},
		}
		infrastructure := client.ASInfrastructure{
			Kubernetes: client.ASCategory{},
			Os:         client.ASCategory{},
		}
		workloads := client.ASWorkloads{
			Kubernetes: client.ASCategory{},
			Os:         client.ASCategory{},
			Cf:         client.ASCategory{},
		}
		category := Category{
			Artifacts:      artifacts,
			Workloads:      workloads,
			Infrastructure: infrastructure,
		}
		as := client.ApplicationScope{
			Name:       "terraform",
			Author:     "administrator",
			OwnerEmail: "email@email.com",
			Categories: category,
		}
		err := cli.UpdateApplicationScope(as)
		if err != nil {
			log.Println("Error updating application_scope: ", err)
		}
	}
}
