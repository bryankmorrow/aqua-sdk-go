package assurance // import "github.com/BryanKMorrow/aqua-sdk-go/types/assurance"

// PermissionList lists globally whitelisted and blacklisted images
type PermissionList struct {
	WhitelistedImages []struct {
		ImageName   string `json:"image_name"`
		Author      string `json:"author"`
		Registry    string `json:"registry"`
		Lastupdated int    `json:"lastupdated"`
		Whitelisted bool   `json:"whitelisted"`
		Blacklisted bool   `json:"blacklisted"`
		Disallowed  bool   `json:"disallowed"`
		Comment     string `json:"comment"`
		Reason      struct {
		} `json:"reason"`
		Pending bool `json:"pending"`
	} `json:"whitelisted_images"`
	BlacklistedImages []struct {
		ImageName   string `json:"image_name"`
		Author      string `json:"author"`
		Registry    string `json:"registry"`
		Lastupdated int    `json:"lastupdated"`
		Whitelisted bool   `json:"whitelisted"`
		Blacklisted bool   `json:"blacklisted"`
		Disallowed  bool   `json:"disallowed"`
		Comment     string `json:"comment"`
		Reason      struct {
		} `json:"reason"`
		Pending bool `json:"pending"`
	} `json:"blacklisted_images"`
}
