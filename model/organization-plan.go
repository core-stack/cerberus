package model

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationPlan struct {
	ID uuid.UUID `json:"id" goqu:"omitempty"`

	OrganizationID *uuid.UUID    `json:"organization_id" validate:"required"`
	Organization   *Organization `json:"organization"`

	PlanID *uuid.UUID `json:"plan_id" validate:"required"`
	Plan   *Plan      `json:"plan"`

	Status PlanStatus `json:"status" validate:"required"`

	CreatedAt time.Time  `json:"created_at" goqu:"omitempty"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
