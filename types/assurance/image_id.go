package assurance // import "github.com/BryanKMorrow/aqua-sdk-go/types/assurance"

type ImageID struct {
	ID          int64  `json:"imageID,omitempty"`
	ImageDigest string `json:"imageDigest,omitempty"`
	ImageName   string `json:"imagename"`
	Author      string `json:"author"`
	Registry    string `json:"registry"`
	LastUpdated int64  `json:"lastupdated,omitempty"`
}
