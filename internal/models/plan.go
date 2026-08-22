package models

type Plan struct {
	FormatVersion string
	TerraformVersion string
	ResourceChanges []ResourceChange
}

type ResourceChange struct {
	Address string
	Mode string
	Type string
	Name string
	Change Change
}