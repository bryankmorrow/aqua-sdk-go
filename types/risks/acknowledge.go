package risks // import "github.com/BryanKMorrow/aqua-sdk-go/types/risks"

import "time"

// Acknowledgements - slice of Ack
type Acknowledgements struct {
	Count    int           `json:"count"`
	Page     int           `json:"page"`
	Pagesize int           `json:"pagesize"`
	Result   []Acknowledge `json:"result"`
}

// Acknowledge stores vulnerabilities whose risk is accepted
type Acknowledge struct {
	IssueType              string    `json:"issue_type"`
	ResourceType           string    `json:"resource_type"`
	ResourceName           string    `json:"resource_name"`
	ResourceVersion        string    `json:"resource_version"`
	ResourcePath           string    `json:"resource_path"`
	ResourceCpe            string    `json:"resource_cpe"`
	IssueName              string    `json:"issue_name"`
	Registry               string    `json:"registry,omitempty"`
	Image                  string    `json:"image,omitempty"`
	Comment                string    `json:"comment"`
	Author                 string    `json:"author"`
	Date                   time.Time `json:"date"`
	FixVersion             string    `json:"fix_version"`
	ExpirationDays         int       `json:"expiration_days"`
	ExpirationConfiguredAt time.Time `json:"expiration_configured_at"`
	ExpirationConfiguredBy string    `json:"expiration_configured_by"`
	Repository             string    `json:"repository,omitempty"`
}
