package main

import (
	"github.com/BryanKMorrow/aqua-sdk-go/client"
	"github.com/BryanKMorrow/aqua-sdk-go/types/assurance"
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
		var variables []assurance.ScopeVariable
		var trusteds []assurance.ImageID
		// Create the variables
		var variable = assurance.ScopeVariable{
			Attribute: "image.name",
			Value:     "*",
		}
		variables = append(variables, variable)
		// Create the Policy Scope
		var scope = assurance.Scope{
			Expression: "v1",
			Variables:  variables,
		}
		// Create the trusted images
		var trusted = assurance.ImageID{
			ImageName: "image:tag",
			Registry:  "registry",
			Author:    "administrator",
		}
		trusteds = append(trusteds, trusted)
		// Create the assurance policy structure
		policy := assurance.Image{
			AssuranceType:            "image",
			Name:                     "aqua-sdk-go",
			Description:              "Example",
			Author:                   "administrator",
			Scope:                    scope,
			CvssSeverity:             "critical",
			CvssSeverityEnabled:      true,
			MaximumScoreEnabled:      true,
			MaximumScore:             8,
			AuditOnFailure:           true,
			TrustedBaseImagesEnabled: true,
			TrustedBaseImages:        trusteds,
			DisallowMalware:          true,
			ScanSensitiveData:        true,
		}
		response := cli.CreateImageAssurance(policy)
		// Print the given installation command for the specified orchestrator
		log.Println(response)
	}
}
