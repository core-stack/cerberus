package model

type PermissionId uint64
type Permission struct {
	ID    PermissionId `json:"id"`
	Group string       `json:"group" validate:"required"`
	Name  string       `json:"name" validate:"required"`
}

const (
	OrganizationUpdate PermissionId = 1 << iota
	OrganizationDelete
	InviteCreate
	InviteUpdate
	InviteDelete
	InviteRead
	MemberCreate
	MemberUpdate
	MemberDelete
	MemberRead
	RoleCreate
	RoleUpdate
	RoleDelete
	RoleRead
)
