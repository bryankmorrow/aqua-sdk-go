package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/BryanKMorrow/aqua-sdk-go/types/assurance"
	"github.com/parnurzeal/gorequest"
)

// GetImageAssuranceName - This returns the Image Assurance Policy by name
// Params: name: The name of the Image Assurance Policy
// Returns: The struct from types/assurance/image
func (cli *Client) GetImageAssuranceName(name string) (assurance.Image, error) {
	var err error
	var response assurance.Image
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/image_assurance/%s", name)
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err = json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetImageAssuranceName from %s%s, %v ", cli.url, apiPath, err.Error())
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	if response.Name == "" {
		err = fmt.Errorf("image assurance policy not found: %s", name)
	}
	return response, err
}

// CreateImageAssurance - Create new Image Assurance Policy
// Post the parameters as a JSON body
// Param: assurance_type: string - Type of assurance policy (image, host, function, cf application)
// Param: name: string - Name of the new policy | !128 character limit
// Param: description: string - Description of the policy
// Param: author: string - Name of user account that created the policy | !Required
// Param: last_update: string - timestamp of last update
// Param: cvss_severity_enabled: bool - scan the cvss severity or not
// Param: cvss_severity: bool - identifier of the cvss severity
// Param: cvss_severity_exclude_no_fix: bool - Indicates that policy should ignore cvss cases that do not have a known fix
// Param: maximum_score_enabled: bool - Indicates if exceeding the maximum score is scanned
// Param: maximum_score: int32 - value of allowed maximum score
// Param: maximum_score_exclude_no_fix: bool - Indicates that policy should ignore cases that do not have a known fix
// Param: custom_checks_enabled: bool - Indicates if scanning should include custom checks | !Doesnt work with Dockerless scanning
// Param: scap_enabled: bool - Indicates if scanning should include scap
// Param: cves_black_list_enabled: bool - Indicates if cves blacklist is relevant
// Param: cves_white_list_enabled: bool - Indicates if cves whitelist is relevant
// Param: packages_black_list_enabled: bool - Enable package blacklist
// Param: packages_white_list_enabled: bool - Enable package whitelist
// Param: only_none_root_users: bool - Only allow non-root users control
// Param: trusted_base_images_enabled: bool - Activate base image control
// Param: scan_sensitive_data: bool - Activate and enable Sensitive data scan
// Param: audit_on_failure: bool - Send audit event if assurance policy fails
// Param: fail_cicd: bool - Send Exit Code to scan results
// Param: block_failed: bool - Block non-compliant images control
// Param: disallow_malware: bool - Block malware control
// Param: blacklist_licenses_enabled: bool - Activate the open source license blacklist control
// Param: blacklisted_licenses: []string - Slice of open source licenses to blacklist
// Param: whitelisted_licenses_enabled: bool - Activate the open source license whitelist control
// Param: whitelisted_licenses: []string - Slice of open source licenses to whitelist
// Param: custom_checks: []assurance.Script - Slice of assurance.Script objects to json
// Param: scap_files: []assurance.Script - Slice of assurance.Script objects to json
// Param: scope: []assurance.Image.Scope - Slice of scope expressions
// Param: registries: []string - Slice of registries
// Param: labels: []string - Slice of labels
// Param: images: []string - Slice of images
// Param: cves_black_list: []string - Slice of cves to blacklist
// Param: cves_white_list: []string - Slice of cves to whitelist
// Param: packages_black_list: []string - Slice of packages to blacklist
// Param: packages_white_list: []string - Slice of packages to whitelist
// Param: allowed_images: []assurance.ImageID - Slice of images by ID
// Param: trusted_base_images: []assurance.ImageID - Slice of images by ID
// Param: readonly: bool - Enable readonly policy permissions
// Param: force_microenforcer: bool - Verify the Entrypoint is microenforcer
// Param: domain: string - Name of the container image
// Param: partial_results_image_fail: bool - return partial results if the image fails assurance check
// Param: control_exclude_no_fix: bool - Enable the No Fix Available exception control
// Param: ignore_recently_published_vln: bool - Enable the Recently published exception
// Param: ignore_recently_published_vln_period: int32 - Number of days to exclude recent vulnerabilities
// Param: ignore_risk_resources_enabled: bool - Ignore specific repositories
// Param: ignored_risk_resources: []string - Slice of ignored resources
// Param: docker_cis_enabled: bool - Enable Docker benchmark for host assurance
// Param: kube_cis_enabled: bool - Enable Kubernetes benchmark for host assurance
// Param: enforce_excessive_permissions: bool - Yeah
// Param: linux_cis_enabled: bool - Enable Linux benchmark for host assurance
// Param: openshift_hardening_enabled: bool - Enable Openshift benchmark for host assurance
// Param: function_integrity_enabled: bool - Enable serverless integrity
func (cli *Client) CreateImageAssurance(policy assurance.Image) string {
	var response = ""
	payload, err := json.Marshal(policy)
	if err != nil {
		log.Println(err)
	}
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/image_assurance")
	resp, _, errs := request.Clone().Post(cli.url + apiPath).Send(string(payload)).End()
	if errs != nil {
		log.Printf("Failed creating assurance policy: %s \n  Status Code: %d", cli.url+apiPath, resp.StatusCode)
	}
	if resp.StatusCode == 204 {
		response = `{"message": "Assurance Policy created successfully"}`
	} else if resp.StatusCode == 400 {
		response = fmt.Sprintf(`{"message": "Assurance Policy creation failed reading body", "status_code": %v}`, resp.StatusCode)
		log.Println(string(payload))
	} else if resp.StatusCode == 500 {
		response = fmt.Sprintf(`{"message": "Assurance Policy creation failed, policy already exists", "status_code": %v}`, resp.StatusCode)
		log.Println(string(payload))
	} else {
		response = fmt.Sprintf(`{"message": "Assurance Policy creation failed", "status_code": %v}`, resp.StatusCode)
		log.Println(string(payload))
	}
	return response
}
