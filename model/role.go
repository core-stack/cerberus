package model

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID uuid.UUID `json:"id" db:"id" goqu:"omitempty"`

	Name    string `json:"name" validate:"required"`
	Enabled bool   `json:"enabled"`

	Permissions PermissionId `json:"permissions" validate:"required"`

	CreatedAt time.Time  `json:"created_at" goqu:"omitempty" validate:"required"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`

	CreatedBy   *User      `json:"created_by"`
	UpdatedBy   *User      `json:"updated_by"`
	DeletedBy   *User      `json:"-"`
	CreatedByID *uuid.UUID `json:"created_by_id"`
	UpdatedByID *uuid.UUID `json:"updated_by_id"`

	OrganizationID *uuid.UUID    `json:"organization_id"`
	Organization   *Organization `json:"organization"`
}
