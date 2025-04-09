package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrCodeNotFound = errors.New("code not found")
	ErrCodeInvalid  = errors.New("code invalid")
	ErrCodeExpired  = errors.New("code expired")
)

type CodeType int

const (
	PasswordReset CodeType = iota
	EmailVerification
)

type Code struct {
	ID int `json:"id" db:"id" goqu:"omitempty"`

	Code string `json:"code" db:"code" validate:"required"`

	Type CodeType `json:"type" db:"type" validate:"required"`

	User   *User     `json:"user" db:"-"`
	UserID uuid.UUID `json:"user_id" db:"user_id" validate:"required"`

	CreatedAt time.Time `json:"created_at" db:"created_at" goqu:"omitempty" validate:"required"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at" validate:"required"`
}

func (c *Code) ValidateCode() error {
	if time.Now().After(c.ExpiresAt) {
		return ErrCodeExpired
	}
	if c.UserID == uuid.Nil {
		return ErrCodeInvalid
	}
	return nil
}

func NewCode(codeType CodeType, userId uuid.UUID) (*Code, error) {
	id := uuid.NewString()

	return &Code{
		Code:      id,
		Type:      codeType,
		ExpiresAt: time.Now().Add(time.Hour * 1),
		UserID:    userId,
	}, nil
}
