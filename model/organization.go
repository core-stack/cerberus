package model

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	ID uuid.UUID `json:"id" goqu:"omitempty"`

	Name        string `json:"name" validate:"required"`
	Slug        string `json:"slug" validate:"required"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
	Enabled     bool   `json:"enabled"`

	Invites []*Invite `json:"invites"`
	Members []*Member `json:"members"`
	Roles   []*Role   `json:"roles"`

	Plan   *OrganizationPlan `json:"plan"`
	PlanID *uuid.UUID        `json:"plan_id"`

	CreatedAt time.Time  `json:"created_at" goqu:"omitempty" validate:"required"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
