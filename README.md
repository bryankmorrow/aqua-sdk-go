# Go client for Aqua Cloud Native Security

This SDK can also be used by your own Go applications to communicate with the Aqua CSP REST API.

For example, to list all the registered images

```go
package main

import (
	"log"

	"github.com/BryanKMorrow/aqua-sdk-go/client"
)

func main() {
	cli := client.NewClient("<AQUA_URL>", "<AQUA_USER>", "<AQUA_PASSWORD>")
	connected := cli.GetAuthToken()
	if connected {
		log.Println("Successfully retrieved JWT Authorization Token")
		registered, remaining, next, total := cli.GetAllImages(1, 14, nil)
        log.Printf("Total: %d  - Remaining: %d - Next Page: %d\n", total, remaining, next)
        for _, image := range registered.Result {
			log.Printf("Found %d vulnerabilities in %s:%s", image.VulnsFound, image.Repository, image.Tag)
		}
	} else {
		log.Fatalln("Failed to retrieve JWT Authorizaiton Token")
	}
}
```
