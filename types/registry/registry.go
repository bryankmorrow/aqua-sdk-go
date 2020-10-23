package registry // import "github.com/BryanKMorrow/aqua-sdk-go/types/registry"

// Registry defines a registry
type Registry []struct {
	Name                     string      `json:"name"`
	Type                     string      `json:"type"`
	DetectedType             int         `json:"detected_type"`
	Description              string      `json:"description"`
	Author                   string      `json:"author"`
	Lastupdate               int         `json:"lastupdate"`
	URL                      string      `json:"url"`
	Username                 string      `json:"-"`
	AutoPull                 bool        `json:"auto_pull"`
	AutoPullTime             string      `json:"auto_pull_time"`
	AutoPullMax              int         `json:"auto_pull_max"`
	PullRepoPatterns         interface{} `json:"pull_repo_patterns"`
	PullRepoPatternsExcluded interface{} `json:"pull_repo_patterns_excluded"`
	PullTagPatterns          interface{} `json:"pull_tag_patterns"`
	PullMaxTags              int         `json:"pull_max_tags"`
	AutoPullRescan           bool        `json:"auto_pull_rescan"`
	Prefixes                 interface{} `json:"prefixes"`
	Webhook                  struct {
		Enabled      bool   `json:"enabled"`
		URL          string `json:"url"`
		AuthToken    string `json:"auth_token"`
		UnQuarantine bool   `json:"un_quarantine"`
	} `json:"webhook"`
	RegistryScanTimeout int           `json:"registry_scan_timeout"`
	PullImageAge        string        `json:"pull_image_age"`
	PullImageTagPattern []interface{} `json:"pull_image_tag_pattern"`
	AlwaysPullPatterns  []interface{} `json:"always_pull_patterns"`
}
