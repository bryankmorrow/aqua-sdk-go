package images  // import "github.com/BryanKMorrow/aqua-sdk-go/types/images"

import "time"

// Sensitive - response from v2/images/registry/repo/tag/sensitive
type Sensitive struct {
	Count    int `json:"count"`
	Page     int `json:"page"`
	Pagesize int `json:"pagesize"`
	Result   []struct {
		Type             string    `json:"type"`
		Path             string    `json:"path"`
		Hash             string    `json:"hash"`
		Filename         string    `json:"filename"`
		Acknowledged     bool      `json:"acknowledged"`
		AcknowledgeDate  time.Time `json:"acknowledge_date"`
		AcknowledgeScope string    `json:"acknowledge_scope"`
	} `json:"result"`
}
