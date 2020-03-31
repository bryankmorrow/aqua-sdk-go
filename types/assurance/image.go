package assurance // import "github.com/BryanKMorrow/aqua-sdk-go/types/assurance"

import "time"

// Image returns a single image struct
type Image struct {
	AssuranceType                    string        `json:"assurance_type"`
	Name                             string        `json:"name"`
	Description                      string        `json:"description,omitempty"`
	Author                           string        `json:"author"`
	Lastupdate                       time.Time     `json:"lastupdate,omitempty"`
	CvssSeverityEnabled              bool          `json:"cvss_severity_enabled,omitempty"`
	CvssSeverity                     string        `json:"cvss_severity,omitempty"`
	CvssSeverityExcludeNoFix         bool          `json:"cvss_severity_exclude_no_fix,omitempty"`
	MaximumScoreEnabled              bool          `json:"maximum_score_enabled,omitempty"`
	MaximumScore                     int           `json:"maximum_score,omitempty"`
	MaximumScoreExcludeNoFix         bool          `json:"maximum_score_exclude_no_fix,omitempty"`
	CustomChecksEnabled              bool          `json:"custom_checks_enabled,omitempty"`
	ScapEnabled                      bool          `json:"scap_enabled,omitempty"`
	CvesBlackListEnabled             bool          `json:"cves_black_list_enabled,omitempty"`
	CvesWhiteListEnabled             bool          `json:"cves_white_list_enabled,omitempty"`
	PackagesBlackListEnabled         bool          `json:"packages_black_list_enabled,omitempty"`
	PackagesWhiteListEnabled         bool          `json:"packages_white_list_enabled,omitempty"`
	OnlyNoneRootUsers                bool          `json:"only_none_root_users,omitempty"`
	TrustedBaseImagesEnabled         bool          `json:"trusted_base_images_enabled,omitempty"`
	ScanSensitiveData                bool          `json:"scan_sensitive_data,omitempty"`
	AuditOnFailure                   bool          `json:"audit_on_failure,omitempty"`
	FailCicd                         bool          `json:"fail_cicd,omitempty"`
	BlockFailed                      bool          `json:"block_failed,omitempty"`
	DisallowMalware                  bool          `json:"disallow_malware,omitempty"`
	BlacklistedLicensesEnabled       bool          `json:"blacklisted_licenses_enabled,omitempty"`
	BlacklistedLicenses              []interface{} `json:"blacklisted_licenses,omitempty"`
	WhitelistedLicensesEnabled       bool          `json:"whitelisted_licenses_enabled,omitempty"`
	WhitelistedLicenses              []interface{} `json:"whitelisted_licenses,omitempty"`
	CustomChecks                     []Scripts     `json:"custom_checks,omitempty"`
	ScapFiles                        []Scripts     `json:"scap_files,omitempty"`
	Scope                            Scope         `json:"scope"`
	Registries                       interface{}   `json:"registries,omitempty"`
	Labels                           interface{}   `json:"labels,omitempty"`
	Images                           interface{}   `json:"images,omitempty"`
	CvesBlackList                    []string      `json:"cves_black_list,omitempty"`
	CvesWhiteList                    []string      `json:"cves_white_list,omitempty"`
	PackagesBlackList                []Package     `json:"packages_black_list,omitempty"`
	PackagesWhiteList                []Package     `json:"packages_white_list,omitempty"`
	AllowedImages                    interface{}   `json:"allowed_images,omitempty"`
	TrustedBaseImages                []ImageID     `json:"trusted_base_images,omitempty"`
	ReadOnly                         bool          `json:"read_only,omitempty"`
	ForceMicroenforcer               bool          `json:"force_microenforcer,omitempty"`
	PartialResultsImageFail          bool          `json:"partial_results_image_fail,omitempty"`
	ControlExcludeNoFix              bool          `json:"control_exclude_no_fix,omitempty"`
	IgnoreRecentlyPublishedVln       bool          `json:"ignore_recently_published_vln,omitempty"`
	IgnoreRecentlyPublishedVlnPeriod int           `json:"ignore_recently_published_vln_period,omitempty"`
	IgnoreRiskResourcesEnabled       bool          `json:"ignore_risk_resources_enabled,omitempty"`
	IgnoredRiskResources             []string      `json:"ignored_risk_resources,omitempty"`
	DockerCisEnabled                 bool          `json:"docker_cis_enabled,omitempty"`
	KubeCisEnabled                   bool          `json:"kube_cis_enabled,omitempty"`
	EnforceExcessivePermissions      bool          `json:"enforce_excessive_permissions,omitempty"`
	LinuxCisEnabled                  bool          `json:"linux_cis_enabled,omitempty"`
	OpenshiftHardeningEnabled        bool          `json:"openshift_hardening_enabled,omitempty"`
	FunctionIntegrityEnabled         bool          `json:"function_integrity_enabled,omitempty"`
}
