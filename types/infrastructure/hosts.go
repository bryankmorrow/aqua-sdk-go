package infrastructure

// Hosts struct
type Hosts struct {
	Count    int `json:"count"`
	Page     int `json:"page"`
	Pagesize int `json:"pagesize"`
	Result   []struct {
		ProjectID          int         `json:"project_id"`
		ID                 string      `json:"id"`
		Logicalname        string      `json:"logicalname"`
		Description        string      `json:"description"`
		Type               string      `json:"type"`
		Version            string      `json:"version"`
		Commit             string      `json:"commit"`
		Hostname           string      `json:"hostname"`
		ShortHostname      string      `json:"short_hostname"`
		Address            string      `json:"address"`
		PublicAddress      string      `json:"public_address"`
		Addresses          interface{} `json:"addresses"`
		Lastupdate         int         `json:"lastupdate"`
		Status             string      `json:"status"`
		Serverid           string      `json:"serverid"`
		ServerName         string      `json:"server_name"`
		DockerVersion      string      `json:"docker_version"`
		HostOs             string      `json:"host_os"`
		IsWindows          bool        `json:"is_windows"`
		HighVulns          int         `json:"high_vulns"`
		MedVulns           int         `json:"med_vulns"`
		LowVulns           int         `json:"low_vulns"`
		NegVulns           int         `json:"neg_vulns"`
		VulnsFound         int         `json:"vulns_found"`
		LastVulnScan       int         `json:"last_vuln_scan"`
		DisplayName        string      `json:"display_name"`
		Compliant          bool        `json:"compliant"`
		Machineid          string      `json:"machineid"`
		ClusterID          int         `json:"cluster_id"`
		Hostlabels         interface{} `json:"hostlabels"`
		Gateways           []string    `json:"gateways"`
		Token              string      `json:"token"`
		Enforce            bool        `json:"enforce"`
		Scan               bool        `json:"scan"`
		TotalPass          int         `json:"total_pass"`
		TotalWarn          int         `json:"total_warn"`
		HostKernelVersion  string      `json:"host_kernel_version"`
		KernelModuleLoaded bool        `json:"kernel_module_loaded"`
		RuntimeProtection  int         `json:"runtime_protection"`
		DockerInfo         struct {
		} `json:"docker_info,omitempty"`
		ContainerActivityProtection bool `json:"container_activity_protection"`
		NetworkProtection           bool `json:"network_protection"`
		UserAccessControl           bool `json:"user_access_control"`
		SyncHostImages              bool `json:"sync_host_images"`
		ImageAssurance              bool `json:"image_assurance"`
		HostProtection              bool `json:"host_protection"`
		Orchestrator                struct {
			Type   string `json:"type"`
			Master bool   `json:"master"`
		} `json:"orchestrator,omitempty"`
		AuditAll          bool        `json:"audit_all"`
		AuditSuccessLogin bool        `json:"audit_success_login"`
		AuditFailedLogin  bool        `json:"audit_failed_login"`
		BatchInstallID    int         `json:"batch_install_id"`
		SaveBatchID       int         `json:"save_batch_id"`
		BatchInstallName  string      `json:"batch_install_name"`
		SyscallEnabled    bool        `json:"syscall_enabled"`
		RuntimeType       string      `json:"runtime_type"`
		InterceptionMode  string      `json:"interception_mode"`
		AquaDigest        string      `json:"aqua_digest"`
		ContainerID       string      `json:"container_id"`
		Secrets           interface{} `json:"secrets"`
	} `json:"result"`
	Query struct {
		IdentifiersOnly bool   `json:"identifiers_only"`
		Status          string `json:"status"`
		ImageName       string `json:"image_name"`
		ImageID         string `json:"image_id"`
		ServerID        string `json:"server_id"`
		BatchName       string `json:"batch_name"`
		Compliant       string `json:"compliant"`
		Address         string `json:"address"`
		Cve             string `json:"cve"`
		ConfigFileName  string `json:"config_file_name"`
	} `json:"query"`
}