package assurance // import "github.com/BryanKMorrow/aqua-sdk-go/types/assurance"

// Package is used for the package_black_list and package_white_list
type Package struct {
	Format       string `json:"format"`
	Name         string `json:"name"`
	Epoch        string `json:"epoch"`
	Version      string `json:"version"`
	VersionRange string `json:"version_range"`
	Release      string `json:"release"`
	Arch         string `json:"arch"`
	License      string `json:"license"`
}
