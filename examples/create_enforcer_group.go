package main

import (
	"github.com/BryanKMorrow/aqua-sdk-go/client"
	"github.com/BryanKMorrow/aqua-sdk-go/types/enforcers"
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
		// Get the gateways
		var gateways = []string{"gatewayID"}
		// Orchestrator options
		var orchestrator = enforcers.Orchestrator{
			Type:           "kubernetes",
			ServiceAccount: "aqua-sa",
			Namespace:      "aqua",
		}
		// Create the enforcer group structure
		enforcer := enforcers.Group{
			AuditFailedLogin:            true,
			AuditSuccessLogin:           true,
			ContainerActivityProtection: true,
			Enforce:                     true,
			Description:                 "This is a test from the aqua-sdk-go package",
			Gateways:                    gateways,
			HostOs:                      "Linux", // Linux and Windows
			Hostname:                    "aqua-sdk-go",
			ID:                          "golang",
			ImageAssurance:              true,
			Logicalname:                 "aqua-sdk-go",
			NetworkProtection:           true,
			Orchestrator:  				 orchestrator,
			RuntimeType:                 "docker", // docker, crio, containerd
			SyncHostImages:              true,
			SyscallEnabled:              false,
			Token:                       "aqua-sdk-token",
			Type:                        "agent", //agent, micro-enforcer, nano-enforcer, vm-enforcer
			UserAccessControl:           true,
			RiskExplorerAutoDiscovery:   false,
			HostProtection:              true,
			HostNetworkProtection:       true,
			EnforcerImageName:           "4.6",
		}
		group := cli.CreateEnforcerGroup(enforcer)
		// Print the given installation command for the specified orchestrator
		log.Println(group.InstallCommand)
	}
}
