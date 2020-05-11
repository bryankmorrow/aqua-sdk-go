package images // import "github.com/BryanKMorrow/aqua-sdk-go/types/images"

// Vulnerabilities - API Response to v2/images/registry/repo/tag/vulnerabilities
type Vulnerabilities struct {
	Count    int             `json:"count"`
	Page     int             `json:"page"`
	Pagesize int             `json:"pagesize"`
	Result   []Vulnerability `json:"result"`
}
