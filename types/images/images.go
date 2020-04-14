package images // import "github.com/BryanKMorrow/aqua-sdk-go/types/images"

// AllResponse - API response to /v2/images
type AllResponse struct {
	Count    int              `json:"count"`
	Page     int              `json:"page"`
	Pagesize int              `json:"pagesize"`
	Result   []SingleResponse `json:"result"`
}
