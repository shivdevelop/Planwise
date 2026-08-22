package parser

type terraformPlan struct {
	FormatVersion string `json:"format_version"`
	TerraformVersion string `json:"terraform_version"`
	ResourceChanges []terraformResourceChange `json:"resource_changes"`
}

type terraformResourceChange struct {
	Address string `json:"address"`
	Mode string `json:"mode"`
	Type string `json:"type"`
	Name string `json:"name"`
	Change terraformChange `json:"change"`
}

type terraformChange struct {
	Actions []string `json:"actions"`
	Before any `json:"before"`
	After any `json:"after"`
	AfterUnknown any `json:"after_unknown"`
}