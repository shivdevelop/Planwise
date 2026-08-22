package models

type Action string

const (
	ActionCreate Action = "create"
	ActionRead Action = "read"
	ActionUpdate Action = "update"
	ActionDelete Action = "delete"
)

type Change struct {
	Actions []Action
	Before any
	After any
	AfterUnknown any
}