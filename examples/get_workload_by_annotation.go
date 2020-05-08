package main

import (
	"encoding/json"
	"github.com/BryanKMorrow/aqua-sdk-go/client"
	"log"
	"os"
	"strings"
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
		cl, _, _, _ := cli.GetContainers(0, 1000, nil)
		for _, container := range cl.Result {
			inspect := cli.InspectContainer(container.ID, container.HostID)
			var label map[string]string
			err := json.Unmarshal(inspect.DockerLabels, &label)
			if err != nil {
				log.Println("error ", err)
				//json: Unmarshal(non-pointer main.Request)
			}
			for k, v := range label {
				if k == "io.kubernetes.pod.namespace" && strings.Contains(v, "sock-shop") {
					total := container.Critical + container.High + container.Medium + container.Low + container.Negligible
					log.Printf("Namespace: %s  Deployment: %s  Pod: %s  Image: %s  Total: %d Critical: %d  High: %d  Malware: %d  SensitiveData: %d\n",
						inspect.KubernetesPodNamespace, inspect.KubernetesPodDeployment, inspect.Name, inspect.Image.Name, total, container.Critical, container.High, container.Malware, container.Sensitive)
					log.Println("Application Team: SockShop")
				}
			}
		}
	}
}
