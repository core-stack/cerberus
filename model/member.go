package model

import (
	"time"

	"github.com/google/uuid"
)

type MemberStatus string

const (
	MemberStatusActive   MemberStatus = "active"
	MemberStatusInactive MemberStatus = "inactive"
	MemberStatusPending  MemberStatus = "pending"
)

type Member struct {
	ID uuid.UUID `json:"id" db:"id" goqu:"omitempty"`

	Email string `json:"email" validate:"required,email"`
	Name  string `json:"name" validate:"required"`

	Organization   *Organization `json:"organization" `
	OrganizationID uuid.UUID     `json:"organization_id" validate:"required"`

	UserID uuid.UUID `json:"user_id" validate:"required"`
	User   *User     `json:"user"`

	RoleID *uuid.UUID `json:"role_id" goqu:"omitempty" validate:"required"`
	Role   *Role      `json:"role"`

	Root   bool         `json:"root" validate:"required"`
	Status MemberStatus `json:"status" validate:"required"`

	CreatedAt time.Time  `json:"created_at" goqu:"omitempty"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
