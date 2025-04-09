package security

import (
	"context"

	"github.com/core-stack/go-security/model"
	"github.com/google/uuid"
)

type UserStore interface {
	FindUserByEmail(ctx context.Context, email string) (*model.User, error)
	FindUserByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	FindUserByNickname(ctx context.Context, name string) (*model.User, error)

	SaveUser(ctx context.Context, user *model.User) error
}
