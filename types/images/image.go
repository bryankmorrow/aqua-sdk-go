package images  // import "github.com/BryanKMorrow/aqua-sdk-go/types/images"

import "time"

// SingleResponse - API response to v2/images/registry/repo/tag
type SingleResponse struct {
	Registry         string    `json:"registry"`
	Name             string    `json:"name"`
	VulnsFound       int       `json:"vulns_found"`
	CritVulns        int       `json:"crit_vulns"`
	HighVulns        int       `json:"high_vulns"`
	MedVulns         int       `json:"med_vulns"`
	LowVulns         int       `json:"low_vulns"`
	NegVulns         int       `json:"neg_vulns"`
	RegistryType     string    `json:"registry_type"`
	Repository       string    `json:"repository"`
	Tag              string    `json:"tag"`
	Created          time.Time `json:"created"`
	Author           string    `json:"author"`
	Digest           string    `json:"digest"`
	Size             int       `json:"size"`
	Os               string    `json:"os"`
	OsVersion        string    `json:"os_version"`
	ScanStatus       string    `json:"scan_status"`
	ScanDate         time.Time `json:"scan_date"`
	ScanError        string    `json:"scan_error"`
	SensitiveData    int       `json:"sensitive_data"`
	Malware          int       `json:"malware"`
	Disallowed       bool      `json:"disallowed"`
	Whitelisted      bool      `json:"whitelisted"`
	Blacklisted      bool      `json:"blacklisted"`
	PermissionAuthor string    `json:"permission_author"`
	PartialResults   bool      `json:"partial_results"`
	NewerImageExists bool      `json:"newer_image_exists"`
	Metadata         struct {
		DockerID      string    `json:"docker_id"`
		Parent        string    `json:"parent"`
		RepoDigests   []string  `json:"repo_digests"`
		Comment       string    `json:"comment"`
		Created       time.Time `json:"created"`
		DockerVersion string    `json:"docker_version"`
		Author        string    `json:"author"`
		Architecture  string    `json:"architecture"`
		Os            string    `json:"os"`
		OsVersion     string    `json:"os_version"`
		Size          int       `json:"size"`
		VirtualSize   int       `json:"virtual_size"`
		DefaultUser   string    `json:"default_user"`
		Env           []string  `json:"env"`
		DockerLabels  []string  `json:"docker_labels"`
		ImageType     string    `json:"image_type"`
	} `json:"metadata"`
	AssuranceResults struct {
		Disallowed bool `json:"disallowed"`
	} `json:"assurance_results"`
	PendingDisallowed     bool `json:"pending_disallowed"`
	MicroenforcerDetected bool `json:"microenforcer_detected"`
}
