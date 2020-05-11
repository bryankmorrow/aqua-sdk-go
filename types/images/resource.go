package images

// Resource is a specific vulnerability resource
type Resource struct {
	Type     string   `json:"type"`
	Format   string   `json:"format"`
	Path     string   `json:"path"`
	Name     string   `json:"name"`
	Version  string   `json:"version"`
	Arch     string   `json:"arch,omitempty"`
	Cpe      string   `json:"cpe"`
	Licenses []string `json:"licenses,omitempty"`
	Hash     string   `json:"hash,omitempty"`
}
