package images // import "github.com/BryanKMorrow/aqua-sdk-go/types/images"

import "time"

// ScanHistory stores the results of the v2/images/{}/{}/{}/scan_history call
type ScanHistory struct {
	Count    int `json:"count"`
	Page     int `json:"page"`
	Pagesize int `json:"pagesize"`
	Result   []struct {
		Date                 time.Time `json:"date"`
		Error                string    `json:"error"`
		Digest               string    `json:"digest"`
		DockerID             string    `json:"docker_id"`
		ImagePulled          bool      `json:"image_pulled"`
		ImageCreationDate    time.Time `json:"image_creation_date"`
		SensitiveDataScanned bool      `json:"sensitive_data_scanned"`
		ExecutablesScanned   bool      `json:"executables_scanned"`
		MalwareScanned       bool      `json:"malware_scanned"`
		CritVulns            int       `json:"crit_vulns"`
		HighVulns            int       `json:"high_vulns"`
		MedVulns             int       `json:"med_vulns"`
		LowVulns             int       `json:"low_vulns"`
		NegVulns             int       `json:"neg_vulns"`
		SensitiveData        int       `json:"sensitive_data"`
		Malware              int       `json:"malware"`
		Disallowed           bool      `json:"disallowed"`
		PartialResults       bool      `json:"partial_results"`
	} `json:"result"`
}
