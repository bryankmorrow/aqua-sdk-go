package assurance // import "github.com/BryanKMorrow/aqua-sdk-go/types/assurance"

// Script is a single user created assurance check
// Just adds snippet field
type Script struct {
	ScriptID     string `json:"script_id"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	LastModified int    `json:"last_modified"`
	Description  string `json:"description"`
	Engine       string `json:"engine"`
	Snippet      string `json:"snippet"`
	ReadOnly     bool   `json:"read_only"`
}
