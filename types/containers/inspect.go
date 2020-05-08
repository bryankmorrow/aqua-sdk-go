package containers

import "encoding/json"

// Inspect holds a running container's metadata
type Inspect struct {
	ID   string   `json:"id"`
	Apps []string `json:"apps"`
	Args []string `json:"args"`
	Host struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"host"`
	Name  string `json:"name"`
	User  string `json:"user"`
	Cforg string `json:"cforg"`
	Image struct {
		ID             string `json:"id"`
		Name           string `json:"name"`
		Registry       string `json:"registry"`
		Registered     bool   `json:"registered"`
		RegisterReason string `json:"register_reason"`
	} `json:"image"`
	Links []interface{} `json:"links"`
	Owner struct {
		ID   string `json:"id"`
		UID  string `json:"uid"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"owner"`
	Level2  string        `json:"level2"`
	Level3  string        `json:"level3"`
	Level4  string        `json:"level4"`
	Memory  int           `json:"memory"`
	Status  string        `json:"status"`
	Devices []interface{} `json:"Devices"`
	Network struct {
		Ports struct {
		} `json:"Ports"`
		Bridge   string `json:"Bridge"`
		Gateway  string `json:"Gateway"`
		Networks struct {
		} `json:"Networks"`
		IPAddress              string      `json:"IPAddress"`
		SandboxID              string      `json:"SandboxID"`
		EndpointID             string      `json:"EndpointID"`
		MacAddress             string      `json:"MacAddress"`
		SandboxKey             string      `json:"SandboxKey"`
		HairpinMode            bool        `json:"HairpinMode"`
		IPPrefixLen            int         `json:"IPPrefixLen"`
		IPv6Gateway            string      `json:"IPv6Gateway"`
		GlobalIPv6Address      string      `json:"GlobalIPv6Address"`
		GlobalIPv6PrefixLen    int         `json:"GlobalIPv6PrefixLen"`
		LinkLocalIPv6Address   string      `json:"LinkLocalIPv6Address"`
		SecondaryIPAddresses   interface{} `json:"SecondaryIPAddresses"`
		LinkLocalIPv6PrefixLen int         `json:"LinkLocalIPv6PrefixLen"`
		SecondaryIPv6Addresses interface{} `json:"SecondaryIPv6Addresses"`
	} `json:"Network"`
	CapAdd  []string      `json:"cap_add"`
	Cfspace string        `json:"cfspace"`
	Command []string      `json:"command"`
	Secrets []interface{} `json:"secrets"`
	Volumes []struct {
		Shared        bool   `json:"shared"`
		HostPath      string `json:"host_path"`
		ReadOnly      bool   `json:"read_only"`
		ContainerPath string `json:"container_path"`
	} `json:"volumes"`
	CapDrop  []string `json:"cap_drop"`
	IpcMode  string   `json:"ipc_mode"`
	PidMode  string   `json:"pid_mode"`
	Services []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"services"`
	UtsMode       string        `json:"uts_mode"`
	Cfappname     string        `json:"cfappname"`
	Cpushares     int           `json:"cpushares"`
	Entrypoint    []string      `json:"entrypoint"`
	Privileged    bool          `json:"privileged"`
	StartTime     int           `json:"start_time"`
	CreateTime    int           `json:"create_time"`
	Description   string        `json:"description"`
	ExtraHosts    []interface{} `json:"extra_hosts"`
	UsernsMode    string        `json:"userns_mode"`
	NetworkMode   string        `json:"network_mode"`
	VolumesFrom   []interface{} `json:"volumes_from"`
	RestartPolicy struct {
		Name              string `json:"Name"`
		MaximumRetryCount int    `json:"MaximumRetryCount"`
	} `json:"RestartPolicy"`
	DockerLabels    json.RawMessage `json:"docker_labels"`
	HealthStatus    string          `json:"health_status"`
	ControllerType  string          `json:"controller_type"`
	ReadonlyRootfs  bool            `json:"readonly_rootfs"`
	ApparmorProfile string          `json:"apparmor_profile"`
	InternalService struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"internal_service"`
	RuntimePolicies      []string      `json:"runtime_policies"`
	RuntimeProfiles      []interface{} `json:"runtime_profiles"`
	SeccompDisabled      bool          `json:"seccomp_disabled"`
	SecurityOptions      []string      `json:"security_options"`
	SystemContainer      bool          `json:"system_container"`
	ModificationTime     int           `json:"modification_time"`
	PublishAllPorts      bool          `json:"publish_all_ports"`
	KubernetesPodName    string        `json:"kubernetes_pod_name"`
	KubernetesPodType    string        `json:"kubernetes_pod_type"`
	EnvironmentVariables []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"environment_variables"`
	KubernetesPodNamespace   string `json:"kubernetes_pod_namespace"`
	KubernetesPodDeployment  string `json:"kubernetes_pod_deployment"`
	KubernetesControllerName string `json:"kubernetes_controller_name"`
	KubernetesControllerType string `json:"kubernetes_controller_type"`
}

type DockerLabel struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
