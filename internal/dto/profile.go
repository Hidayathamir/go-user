package dto

import (
	"errors"
	"time"

	"github.com/Hidayathamir/go-user/internal/entity"
)

// ResGetProfileByUsername -.
type ResGetProfileByUsername struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// LoadEntityUser load from entity.User then return ResGetProfileByUsername.
func (r ResGetProfileByUsername) LoadEntityUser(user entity.User) ResGetProfileByUsername {
	return ResGetProfileByUsername{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// ReqUpdateProfileByUserID -.
type ReqUpdateProfileByUserID struct {
	UserJWT  string `json:"-"`
	Password string `json:"password"`
}

// Validate validate ReqUpdateProfileByUserID.
func (r ReqUpdateProfileByUserID) Validate() error {
	if r.UserJWT == "" {
		return errors.New("ReqUpdateProfileByUserID.UserJWT can not be empty")
	}
	return nil
}

// ToEntityUser transform ReqUpdateProfileByUserID to entity.User.
func (r ReqUpdateProfileByUserID) ToEntityUser() entity.User {
	return entity.User{
		Password: r.Password,
	}
}
