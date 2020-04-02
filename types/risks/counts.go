package risks // import "github.com/BryanKMorrow/aqua-sdk-go/types/risks"

// Counts contains the risk information
type Counts struct {
	RegisteredImages        int     `json:"registered_images"`
	VulnerableImages        int     `json:"vulnerable_images"`
	VulnerableImagesPercent float64 `json:"vulnerable_images_percent"`
	TotalVulns              int     `json:"total_vulns"`
	HighVulns               int     `json:"high_vulns"`
	HighVulnsPercent        float64 `json:"high_vulns_percent"`
	MedVulns                int     `json:"med_vulns"`
	MedVulnsPercent         float64 `json:"med_vulns_percent"`
	LowVulns                int     `json:"low_vulns"`
	LowVulnsPercent         float64 `json:"low_vulns_percent"`
	NegVulns                int     `json:"neg_vulns"`
	NegVulnsPercent         float64 `json:"neg_vulns_percent"`
}
