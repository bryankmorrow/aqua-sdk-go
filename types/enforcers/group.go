package enforcers // import "github.com/BryanKMorrow/aqua-sdk-go/types/enforcers"

// Orchestrator is used to inject into the Group struct
type Orchestrator struct {
	Type           string `json:"type"`
	ServiceAccount string `json:"service_account"`
	Namespace      string `json:"namespace"`
}

// Group is the request and response format for an Enforcer Group (hostbatch)
type Group struct {
	AllowedLabels               []interface{} `json:"allowed_labels"`
	AllowedRegistries           []interface{} `json:"allowed_registries"`
	AllowedSecrets              []interface{} `json:"allowed_secrets"`
	AuditFailedLogin            bool          `json:"audit_failed_login"`
	AuditSuccessLogin           bool          `json:"audit_success_login"`
	ContainerActivityProtection bool          `json:"container_activity_protection"`
	Enforce                     bool          `json:"enforce"`
	Description                 string        `json:"description"`
	Gateways                    []string      `json:"gateways"`
	HostOs                      string        `json:"host_os"`
	Hostname                    string        `json:"hostname"`
	ID                          string        `json:"id"`
	ImageAssurance              bool          `json:"image_assurance"`
	Logicalname                 string        `json:"logicalname"`
	NetworkProtection           bool          `json:"network_protection"`
	Orchestrator                struct {
		Type           string `json:"type"`
		ServiceAccount string `json:"service_account"`
		Namespace      string `json:"namespace"`
	} `json:"orchestrator"`
	RuntimeType               string        `json:"runtime_type"`
	SyncHostImages            bool          `json:"sync_host_images"`
	SyscallEnabled            bool          `json:"syscall_enabled"`
	Token                     string        `json:"token"`
	Type                      string        `json:"type"`
	UserAccessControl         bool          `json:"user_access_control"`
	AllowedLabelsTemp         []interface{} `json:"allowed_labels_temp"`
	AllowedRegistriesTemp     []interface{} `json:"allowed_registries_temp"`
	AllowedSecretsTemp        []interface{} `json:"allowed_secrets_temp"`
	RiskExplorerAutoDiscovery bool          `json:"risk_explorer_auto_discovery"`
	RuntimePolicyName         string        `json:"runtime_policy_name"`
	HostProtection            bool          `json:"host_protection"`
	HostNetworkProtection     bool          `json:"host_network_protection"`
	EnforcerImageName         string        `json:"enforcer_image_name"`
}

// GroupResponse is the return after group creation
type GroupResponse struct {
	ID                          string   `json:"id"`
	Logicalname                 string   `json:"logicalname"`
	Type                        string   `json:"type"`
	EnforcerImageName           string   `json:"enforcer_image_name"`
	Description                 string   `json:"description"`
	Gateways                    []string `json:"gateways"`
	GatewayName                 string   `json:"gateway_name"`
	GatewayAddress              string   `json:"gateway_address"`
	Enforce                     bool     `json:"enforce"`
	ContainerActivityProtection bool     `json:"container_activity_protection"`
	NetworkProtection           bool     `json:"network_protection"`
	HostNetworkProtection       bool     `json:"host_network_protection"`
	UserAccessControl           bool     `json:"user_access_control"`
	ImageAssurance              bool     `json:"image_assurance"`
	HostProtection              bool     `json:"host_protection"`
	AuditAll                    bool     `json:"audit_all"`
	AuditSuccessLogin           bool     `json:"audit_success_login"`
	AuditFailedLogin            bool     `json:"audit_failed_login"`
	LastUpdate                  int      `json:"last_update"`
	Token                       string   `json:"token"`
	Command                     struct {
		Default    string `json:"default"`
		Kubernetes string `json:"kubernetes"`
		Swarm      string `json:"swarm"`
		Windows    string `json:"windows"`
	} `json:"command"`
	Orchestrator struct {
		Type           string `json:"type"`
		Master         bool   `json:"master"`
		ServiceAccount string `json:"service_account"`
		Namespace      string `json:"namespace"`
	} `json:"orchestrator"`
	HostOs                    string `json:"host_os"`
	InstallCommand            string `json:"install_command"`
	HostsCount                int    `json:"hosts_count"`
	DisconnectedCount         int    `json:"disconnected_count"`
	ConnectedCount            int    `json:"connected_count"`
	HighVulns                 int    `json:"high_vulns"`
	MedVulns                  int    `json:"med_vulns"`
	LowVulns                  int    `json:"low_vulns"`
	NegVulns                  int    `json:"neg_vulns"`
	SyscallEnabled            bool   `json:"syscall_enabled"`
	RuntimeType               string `json:"runtime_type"`
	SyncHostImages            bool   `json:"sync_host_images"`
	RiskExplorerAutoDiscovery bool   `json:"risk_explorer_auto_discovery"`
	RuntimePolicyName         string `json:"runtime_policy_name"`
	PasDeploymentLink         string `json:"pas_deployment_link"`
	AquaVersion               string `json:"aqua_version"`
}
