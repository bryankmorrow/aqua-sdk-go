package images

import "time"

// AllResponse - API response to /v2/images
type AllResponse struct {
	Count    int `json:"count"`
	Page     int `json:"page"`
	Pagesize int `json:"pagesize"`
	Result   []struct {
		Registry      string      `json:"registry"`
		Name          string      `json:"name"`
		VulnsFound    int         `json:"vulns_found"`
		CritVulns     int         `json:"crit_vulns"`
		HighVulns     int         `json:"high_vulns"`
		MedVulns      int         `json:"med_vulns"`
		LowVulns      int         `json:"low_vulns"`
		NegVulns      int         `json:"neg_vulns"`
		Repository    string      `json:"repository"`
		Tag           string      `json:"tag"`
		Created       time.Time   `json:"created"`
		Author        string      `json:"author"`
		Digest        string      `json:"digest"`
		Size          int         `json:"size"`
		Labels        interface{} `json:"labels"`
		Os            string      `json:"os"`
		OsVersion     string      `json:"os_version"`
		ScanStatus    string      `json:"scan_status"`
		ScanDate      time.Time   `json:"scan_date"`
		ScanError     string      `json:"scan_error"`
		SensitiveData int         `json:"sensitive_data"`
		Malware       int         `json:"malware"`
		Disallowed    bool        `json:"disallowed"`
		Whitelisted   bool        `json:"whitelisted"`
		Blacklisted   bool        `json:"blacklisted"`
		// PolicyFailures   []interface{} `json:"policy_failures"`
		PartialResults   bool `json:"partial_results"`
		NewerImageExists bool `json:"newer_image_exists"`
		AssuranceResults struct {
			Disallowed      bool        `json:"disallowed"`
			ChecksPerformed interface{} `json:"checks_performed"`
		} `json:"assurance_results"`
		PendingDisallowed     bool `json:"pending_disallowed"`
		MicroenforcerDetected bool `json:"microenforcer_detected"`
	} `json:"result"`
}
