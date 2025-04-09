package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type UserChallenge string

func (u UserChallenge) Value() (driver.Value, error) {
	return string(u), nil
}

func (u *UserChallenge) Scan(value any) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to scan UserChallenge: %v", value)
	}
	*u = UserChallenge(str)
	return nil
}

// Para o slice
type UserChallengeList []UserChallenge

func (u UserChallengeList) Value() (driver.Value, error) {
	return json.Marshal(u)
}

func (u *UserChallengeList) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan UserChallengeList: %v", value)
	}
	return json.Unmarshal(bytes, u)
}

const (
	UserChallengeActiveAccount  UserChallenge = "active_account"
	UserChallengeResetPassword  UserChallenge = "reset_password"
	UserChallengeChangePassword UserChallenge = "change_password"
	UserChallengeInvite         UserChallenge = "invite"
)

type UserStatus int

const (
	UserStatusPending UserStatus = iota
	UserStatusActive
	UserStatusBanned
	UserStatusTemporaryBanned
	UserStatusDeleted
)

type User struct {
	ID        uuid.UUID       `json:"id"`
	Name      string          `json:"name"`
	Nickname  *string         `json:"nickname"`
	Email     string          `json:"email"`
	Password  string          `json:"-"`
	Status    UserStatus      `json:"status"`
	Challenge []UserChallenge `json:"challenge"`

	RoleID *int  `json:"role_id"`
	Role   *Role `json:"role"`

	OrganizationID *int64        `json:"organization_id"`
	Organization   *Organization `json:"organization"`

	Members []*Member `json:"members"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
