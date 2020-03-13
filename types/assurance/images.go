package assurance // import "github.com/BryanKMorrow/aqua-sdk-go/types/assurance"

import "time"


// Images - Assurance Policy list  from v2/image_assurance
type Images struct {
	Count    int `json:"count"`
	Page     int `json:"page"`
	Pagesize int `json:"pagesize"`
	Result   []struct {
		AssuranceType              string      `json:"assurance_type"`
		Name                       string      `json:"name"`
		Description                string      `json:"description"`
		Author                     string      `json:"author"`
		Lastupdate                 time.Time   `json:"lastupdate"`
		CvssSeverityEnabled        bool        `json:"cvss_severity_enabled"`
		CvssSeverity               string      `json:"cvss_severity"`
		CvssSeverityExcludeNoFix   bool        `json:"cvss_severity_exclude_no_fix"`
		MaximumScoreEnabled        bool        `json:"maximum_score_enabled"`
		MaximumScore               int         `json:"maximum_score"`
		MaximumScoreExcludeNoFix   bool        `json:"maximum_score_exclude_no_fix"`
		CustomChecksEnabled        bool        `json:"custom_checks_enabled"`
		ScapEnabled                bool        `json:"scap_enabled"`
		CvesBlackListEnabled       bool        `json:"cves_black_list_enabled"`
		CvesWhiteListEnabled       bool        `json:"cves_white_list_enabled"`
		PackagesBlackListEnabled   bool        `json:"packages_black_list_enabled"`
		PackagesWhiteListEnabled   bool        `json:"packages_white_list_enabled"`
		OnlyNoneRootUsers          bool        `json:"only_none_root_users"`
		TrustedBaseImagesEnabled   bool        `json:"trusted_base_images_enabled"`
		ScanSensitiveData          bool        `json:"scan_sensitive_data"`
		AuditOnFailure             bool        `json:"audit_on_failure"`
		FailCicd                   bool        `json:"fail_cicd"`
		BlockFailed                bool        `json:"block_failed"`
		DisallowMalware            bool        `json:"disallow_malware"`
		BlacklistedLicensesEnabled bool        `json:"blacklisted_licenses_enabled"`
		BlacklistedLicenses        interface{} `json:"blacklisted_licenses"`
		WhitelistedLicensesEnabled bool        `json:"whitelisted_licenses_enabled"`
		WhitelistedLicenses        interface{} `json:"whitelisted_licenses"`
		CustomChecks               interface{} `json:"custom_checks"`
		ScapFiles                  interface{} `json:"scap_files"`
		Scope                      struct {
			Expression string `json:"expression"`
			Variables  []struct {
				Attribute string `json:"attribute"`
				Value     string `json:"value"`
			} `json:"variables"`
		} `json:"scope"`
		Registries                       interface{} `json:"registries"`
		Labels                           interface{} `json:"labels"`
		Images                           interface{} `json:"images"`
		CvesBlackList                    []string    `json:"cves_black_list"`
		CvesWhiteList                    []string    `json:"cves_white_list"`
		PackagesBlackList                interface{} `json:"packages_black_list"`
		PackagesWhiteList                interface{} `json:"packages_white_list"`
		AllowedImages                    interface{} `json:"allowed_images"`
		TrustedBaseImages                interface{} `json:"trusted_base_images"`
		ReadOnly                         bool        `json:"read_only"`
		ForceMicroenforcer               bool        `json:"force_microenforcer"`
		PartialResultsImageFail          bool        `json:"partial_results_image_fail"`
		ControlExcludeNoFix              bool        `json:"control_exclude_no_fix"`
		IgnoreRecentlyPublishedVln       bool        `json:"ignore_recently_published_vln"`
		IgnoreRecentlyPublishedVlnPeriod int         `json:"ignore_recently_published_vln_period"`
		IgnoreRiskResourcesEnabled       bool        `json:"ignore_risk_resources_enabled"`
		IgnoredRiskResources             []string    `json:"ignored_risk_resources"`
		DockerCisEnabled                 bool        `json:"docker_cis_enabled"`
		KubeCisEnabled                   bool        `json:"kube_cis_enabled"`
		EnforceExcessivePermissions      bool        `json:"enforce_excessive_permissions"`
		LinuxCisEnabled                  bool        `json:"linux_cis_enabled"`
		OpenshiftHardeningEnabled        bool        `json:"openshift_hardening_enabled"`
		FunctionIntegrityEnabled         bool        `json:"function_integrity_enabled"`
	} `json:"result"`
}

// Image returns a single image struct
type Image struct {
	AssuranceType              string        `json:"assurance_type"`
	Name                       string        `json:"name"`
	Description                string        `json:"description"`
	Author                     string        `json:"author"`
	Lastupdate                 time.Time     `json:"lastupdate"`
	CvssSeverityEnabled        bool          `json:"cvss_severity_enabled"`
	CvssSeverity               string        `json:"cvss_severity"`
	CvssSeverityExcludeNoFix   bool          `json:"cvss_severity_exclude_no_fix"`
	MaximumScoreEnabled        bool          `json:"maximum_score_enabled"`
	MaximumScore               int           `json:"maximum_score"`
	MaximumScoreExcludeNoFix   bool          `json:"maximum_score_exclude_no_fix"`
	CustomChecksEnabled        bool          `json:"custom_checks_enabled"`
	ScapEnabled                bool          `json:"scap_enabled"`
	CvesBlackListEnabled       bool          `json:"cves_black_list_enabled"`
	CvesWhiteListEnabled       bool          `json:"cves_white_list_enabled"`
	PackagesBlackListEnabled   bool          `json:"packages_black_list_enabled"`
	PackagesWhiteListEnabled   bool          `json:"packages_white_list_enabled"`
	OnlyNoneRootUsers          bool          `json:"only_none_root_users"`
	TrustedBaseImagesEnabled   bool          `json:"trusted_base_images_enabled"`
	ScanSensitiveData          bool          `json:"scan_sensitive_data"`
	AuditOnFailure             bool          `json:"audit_on_failure"`
	FailCicd                   bool          `json:"fail_cicd"`
	BlockFailed                bool          `json:"block_failed"`
	DisallowMalware            bool          `json:"disallow_malware"`
	BlacklistedLicensesEnabled bool          `json:"blacklisted_licenses_enabled"`
	BlacklistedLicenses        []interface{} `json:"blacklisted_licenses"`
	WhitelistedLicensesEnabled bool          `json:"whitelisted_licenses_enabled"`
	WhitelistedLicenses        []interface{} `json:"whitelisted_licenses"`
	CustomChecks               []interface{} `json:"custom_checks"`
	ScapFiles                  []interface{} `json:"scap_files"`
	Scope                      struct {
		Expression string `json:"expression"`
		Variables  []struct {
			Attribute string `json:"attribute"`
			Value     string `json:"value"`
		} `json:"variables"`
	} `json:"scope"`
	Registries                       interface{} `json:"registries"`
	Labels                           interface{} `json:"labels"`
	Images                           interface{} `json:"images"`
	CvesBlackList                    []string    `json:"cves_black_list"`
	CvesWhiteList                    []string    `json:"cves_white_list"`
	PackagesBlackList                interface{} `json:"packages_black_list"`
	PackagesWhiteList                interface{} `json:"packages_white_list"`
	AllowedImages                    interface{} `json:"allowed_images"`
	TrustedBaseImages                interface{} `json:"trusted_base_images"`
	ReadOnly                         bool        `json:"read_only"`
	ForceMicroenforcer               bool        `json:"force_microenforcer"`
	PartialResultsImageFail          bool        `json:"partial_results_image_fail"`
	ControlExcludeNoFix              bool        `json:"control_exclude_no_fix"`
	IgnoreRecentlyPublishedVln       bool        `json:"ignore_recently_published_vln"`
	IgnoreRecentlyPublishedVlnPeriod int         `json:"ignore_recently_published_vln_period"`
	IgnoreRiskResourcesEnabled       bool        `json:"ignore_risk_resources_enabled"`
	IgnoredRiskResources             []string    `json:"ignored_risk_resources"`
	DockerCisEnabled                 bool        `json:"docker_cis_enabled"`
	KubeCisEnabled                   bool        `json:"kube_cis_enabled"`
	EnforceExcessivePermissions      bool        `json:"enforce_excessive_permissions"`
	LinuxCisEnabled                  bool        `json:"linux_cis_enabled"`
	OpenshiftHardeningEnabled        bool        `json:"openshift_hardening_enabled"`
	FunctionIntegrityEnabled         bool        `json:"function_integrity_enabled"`
}
