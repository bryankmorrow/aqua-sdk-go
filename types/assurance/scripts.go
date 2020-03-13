package assurance // import "github.com/BryanKMorrow/aqua-sdk-go/types/assurance"

// Scripts are user created assurance checks
type Scripts []struct {
	ScriptID     string `json:"script_id"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	LastModified int    `json:"last_modified"`
	Description  string `json:"description"`
	Engine       string `json:"engine"`
	ReadOnly     bool   `json:"read_only"`
}
