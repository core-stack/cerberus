package model

import (
	"time"

	"github.com/google/uuid"
)

type Invite struct {
	ID uuid.UUID `json:"id" db:"id" goqu:"omitempty"`

	User   *User      `json:"user"`
	UserID *uuid.UUID `json:"user_id" validate:"required"`

	Organization   *Organization `json:"organization"`
	OrganizationID uuid.UUID     `json:"organization_id" validate:"required"`

	Email string `json:"email" validate:"required"`

	Token uuid.UUID `json:"-" goqu:"omitempty" validate:"required"`

	ExpiresAt time.Time `json:"expires_at" validate:"required"`

	CreatedAt time.Time  `json:"created_at" goqu:"omitempty" validate:"required"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`

	CreatedBy   *Member `json:"created_by"`
	CreatedByID *int64  `json:"created_by_id" validate:"required"`
}
