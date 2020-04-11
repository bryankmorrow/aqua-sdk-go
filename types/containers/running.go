package containers

// Containers holds the result of /v1/containers
type Containers struct {
	Count    int `json:"count"`
	Page     int `json:"page"`
	Pagesize int `json:"pagesize"`
	Result   []struct {
		ID                 string      `json:"id"`
		HostID             string      `json:"host_id"`
		Name               string      `json:"name"`
		HaveSecrets        bool        `json:"have_secrets"`
		HostName           string      `json:"host_name"`
		HostLname          string      `json:"host_lname"`
		ImageDigest        string      `json:"image_digest"`
		ServerDigest       string      `json:"server_digest"`
		ImageName          string      `json:"image_name"`
		OriginImageName    string      `json:"origin_image_name"`
		OwnerName          string      `json:"owner_name"`
		ImageID            string      `json:"image_id"`
		Status             string      `json:"status"`
		RiskLevel          string      `json:"risk_level"`
		IsEvaluated        bool        `json:"is_evaluated"`
		ValidDigest        bool        `json:"valid_digest"`
		IsProfiling        bool        `json:"is_profiling"`
		IsRegistered       bool        `json:"is_registered"`
		IsDisallowed       bool        `json:"is_disallowed"`
		IsRoot             bool        `json:"is_root"`
		IsPrivileged       bool        `json:"is_privileged"`
		ScanStatus         string      `json:"scan_status"`
		CreateTime         int         `json:"create_time"`
		SystemContainer    bool        `json:"system_container"`
		StartTime          int         `json:"start_time"`
		ModifyTime         int         `json:"modify_time"`
		AquaService        string      `json:"aqua_service"`
		Secrets            interface{} `json:"secrets"`
		Risk               int         `json:"risk"`
		NetworkMode        string      `json:"network_mode"`
		RegistryImageName  string      `json:"registry_image_name"`
		RuntimeProfileName string      `json:"runtime_profile_name"`
		ContainerType      string      `json:"container_type"`
		Total              int         `json:"total"`
		Critical           int         `json:"critical"`
		High               int         `json:"high"`
		Medium             int         `json:"medium"`
		Low                int         `json:"low"`
		Sensitive          int         `json:"sensitive"`
		Malware            int         `json:"malware"`
		Negligible         int         `json:"negligible"`
	} `json:"result"`
}
