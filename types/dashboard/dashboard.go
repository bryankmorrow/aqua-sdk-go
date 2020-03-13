package dashboard // import "github.com/BryanKMorrow/aqua-sdk-go/types/dashboard

// Overview is the response from v1/dashboard
type Overview struct {
	Filter struct {
		Application string `json:"application"`
		Registry    string `json:"registry"`
		Hosts       string `json:"hosts"`
	} `json:"filter"`
	RunningContainers struct {
		Total        int `json:"total"`
		Critical     int `json:"critical"`
		High         int `json:"high"`
		Medium       int `json:"medium"`
		Low          int `json:"low"`
		Ok           int `json:"ok"`
		Unregistered int `json:"unregistered"`
	} `json:"running_containers"`
	RegistryCounts struct {
		Images struct {
			Total    int `json:"total"`
			Critical int `json:"critical"`
			High     int `json:"high"`
			Medium   int `json:"medium"`
			Low      int `json:"low"`
			Ok       int `json:"ok"`
		} `json:"images"`
		ImagesTrends    interface{} `json:"images_trends"`
		Vulnerabilities struct {
			Total    int `json:"total"`
			Critical int `json:"critical"`
			High     int `json:"high"`
			Medium   int `json:"medium"`
			Low      int `json:"low"`
			Ok       int `json:"ok"`
		} `json:"vulnerabilities"`
		CvesTrends interface{} `json:"cves_trends"`
	} `json:"registry_counts"`
	Hosts struct {
		Total             int         `json:"total"`
		DisconnectedCount int         `json:"disconnected_count"`
		Hosts             interface{} `json:"hosts"`
	} `json:"hosts"`
	Alerts []struct {
		ID           int    `json:"id"`
		Time         int    `json:"time"`
		Type         string `json:"type"`
		User         string `json:"user"`
		Image        string `json:"image"`
		Imagehash    string `json:"imagehash"`
		Container    string `json:"container"`
		Containerid  string `json:"containerid"`
		Host         string `json:"host"`
		Hostid       string `json:"hostid"`
		Category     string `json:"category"`
		Result       int    `json:"result"`
		UserResponse string `json:"user_response"`
		Data         string `json:"data"`
	} `json:"alerts"`
	AuditTickers []struct {
		ID          int    `json:"id"`
		Time        int    `json:"time"`
		Type        string `json:"type"`
		User        string `json:"user"`
		Action      string `json:"action"`
		Image       string `json:"image"`
		Imagehash   string `json:"imagehash"`
		Container   string `json:"container"`
		Containerid string `json:"containerid"`
		Host        string `json:"host"`
		Hostid      string `json:"hostid"`
		Category    string `json:"category"`
		Result      int    `json:"result"`
		Data        string `json:"data"`
	} `json:"audit_tickers"`
}
