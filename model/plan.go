package model

import (
	"time"

	"github.com/google/uuid"
)

type PlanStatus string

const (
	PlanStatusActive   PlanStatus = "active"
	PlanStatusInactive PlanStatus = "inactive"
	PlanStatusPending  PlanStatus = "pending"
)

type Plan struct {
	ID uuid.UUID `json:"id" db:"id" goqu:"omitempty"`

	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int    `json:"price" validate:"required"`

	CreatedAt time.Time  `json:"created_at" goqu:"omitempty" validate:"required"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}
