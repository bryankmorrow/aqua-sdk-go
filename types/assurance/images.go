package assurance // import "github.com/BryanKMorrow/aqua-sdk-go/types/assurance"

// Images - Assurance Policy list  from v2/image_assurance
type Images struct {
	Count    int     `json:"count"`
	Page     int     `json:"page"`
	Pagesize int     `json:"pagesize"`
	Result   []Image `json:"result"`
}
