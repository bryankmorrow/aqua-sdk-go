package infrastructure  // import "github.com/BryanKMorrow/aqua-sdk-go/types/infrastructure

import "time"

// List returns all known infrastructure - v2/infrastructure
// INT params - page and page_size
// STRING params - order_by  (name|type|author), search and type (nodes|clusters)
// BOOL params - enforced
type List struct {
	Count    int `json:"count"`
	Page     int `json:"page"`
	Pagesize int `json:"pagesize"`
	Result   []struct {
		ID                     int    `json:"id"`
		ClusterID              int    `json:"cluster_id"`
		NodeID                 string `json:"node_id"`
		Name                   string `json:"name"`
		Type                   string `json:"type"`
		ClusterVulnerabilities struct {
			Lastupdate     int `json:"lastupdate"`
			Kubehunterscan struct {
				Ok    int    `json:"ok"`
				Low   int    `json:"low"`
				High  int    `json:"high"`
				Kburl string `json:"kburl"`
				Nodes []struct {
					Type     string `json:"type"`
					Location string `json:"location"`
				} `json:"nodes"`
				Medium   int `json:"medium"`
				Services []struct {
					Type     string `json:"type"`
					Service  string `json:"service"`
					Location string `json:"location"`
				} `json:"services"`
				Vulnerabilities []struct {
					URL           string `json:"url"`
					Vid           string `json:"vid"`
					Hunter        string `json:"hunter"`
					Category      string `json:"category"`
					Evidence      string `json:"evidence"`
					Location      string `json:"location"`
					Severity      string `json:"severity"`
					Description   string `json:"description"`
					Vulnerability string `json:"vulnerability"`
				} `json:"vulnerabilities"`
				HunterStatistics []struct {
					Name            string `json:"name"`
					Description     string `json:"description"`
					Vulnerabilities int    `json:"vulnerabilities"`
				} `json:"hunter_statistics"`
			} `json:"kubehunterscan"`
		} `json:"cluster_vulnerabilities,omitempty"`
		Data struct {
			Roles struct {
				ClusterRoles []struct {
					UID   string `json:"uid"`
					Name  string `json:"name"`
					Rules []struct {
						Verbs     []string `json:"verbs"`
						Apigroups []string `json:"apigroups"`
						Resources []string `json:"resources"`
					} `json:"rules"`
					Labels          []string `json:"labels,omitempty"`
					Creationtime    string   `json:"creationtime"`
					Resourcecount   int      `json:"resourcecount,omitempty"`
					Resourceversion string   `json:"resourceversion"`
				} `json:"cluster_roles"`
			} `json:"roles"`
			Information struct {
				ServerInfo struct {
					Major      string    `json:"major"`
					Minor      string    `json:"minor"`
					Platform   string    `json:"platform"`
					Builddate  time.Time `json:"builddate"`
					Gitversion string    `json:"gitversion"`
				} `json:"server_info"`
				ComponentsStatusList []struct {
					Name       string `json:"name"`
					Conditions []struct {
						Type    string `json:"type"`
						Status  string `json:"status"`
						Message string `json:"message"`
					} `json:"conditions"`
					Creationtime string `json:"creationtime"`
				} `json:"components_status_list"`
			} `json:"information"`
		} `json:"data,omitempty"`
		SecurityIssues struct {
			CritVulns    int `json:"crit_vulns"`
			HighVulns    int `json:"high_vulns"`
			MedVulns     int `json:"med_vulns"`
			LowVulns     int `json:"low_vulns"`
			NegVulns     int `json:"neg_vulns"`
			Malware      int `json:"malware"`
			LastVulnScan int `json:"last_vuln_scan"`
		} `json:"security_issues"`
		IsEnforced             bool      `json:"is_enforced"`
		CreatedDate            time.Time `json:"created_date"`
	} `json:"result"`
	Query struct {
		IdentifiersOnly bool   `json:"identifiers_only"`
		Type            string `json:"type"`
		Enforced        string `json:"enforced"`
	} `json:"query"`
}
