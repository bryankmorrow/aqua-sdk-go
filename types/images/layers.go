package images  // import "github.com/BryanKMorrow/aqua-sdk-go/types/images"

// Layers response from v2/images/registry/repo/tag/history_layers
type Layers struct {
	Count    int `json:"count"`
	Page     int `json:"page"`
	Pagesize int `json:"pagesize"`
	Result   []struct {
		ID            string `json:"id"`
		Comment       string `json:"comment"`
		Created       int    `json:"created"`
		CreatedBy     string `json:"created_by"`
		Size          int    `json:"size"`
		Position      int    `json:"position"`
		CritVulns     int    `json:"crit_vulns"`
		HighVulns     int    `json:"high_vulns"`
		MedVulns      int    `json:"med_vulns"`
		LowVulns      int    `json:"low_vulns"`
		NegVulns      int    `json:"neg_vulns"`
		SensitiveData int    `json:"sensitive_data"`
		VulnsFound    int    `json:"vulns_found"`
		IsBaseLayer   bool   `json:"is_base_layer"`
	} `json:"result"`
}

// LayerResult - an attempt to de-duplicate layer history
type LayerResult []struct {
	ID            string `json:"id"`
	Comment       string `json:"comment"`
	Created       int    `json:"created"`
	CreatedBy     string `json:"created_by"`
	Size          int    `json:"size"`
	Position      int    `json:"position"`
	CritVulns     int    `json:"crit_vulns"`
	HighVulns     int    `json:"high_vulns"`
	MedVulns      int    `json:"med_vulns"`
	LowVulns      int    `json:"low_vulns"`
	NegVulns      int    `json:"neg_vulns"`
	SensitiveData int    `json:"sensitive_data"`
	VulnsFound    int    `json:"vulns_found"`
	IsBaseLayer   bool   `json:"is_base_layer"`
}
