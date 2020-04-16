package images // import "github.com/BryanKMorrow/aqua-sdk-go/types/images"

// Images - API response to /v2/images
type Images struct {
	Count    int     `json:"count"`
	Page     int     `json:"page"`
	Pagesize int     `json:"pagesize"`
	Result   []Image `json:"result"`
}
