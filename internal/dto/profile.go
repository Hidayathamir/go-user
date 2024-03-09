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
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// LoadEntityUser load from entity.User then return ResGetProfileByUsername.
func (r ResGetProfileByUsername) LoadEntityUser(user entity.User) ResGetProfileByUsername {
	return ResGetProfileByUsername{
		ID:        user.ID,
		Username:  user.Username,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// ReqUpdateProfileByUsername -.
type ReqUpdateProfileByUsername struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate validate ReqUpdateProfileByUsername.
func (r ReqUpdateProfileByUsername) Validate() error {
	if r.Username == "" {
		return errors.New("ReqUpdateProfileByUsername.Username can not be empty")
	}
	if r.Password == "" {
		return errors.New("ReqUpdateProfileByUsername.Password can not be empty")
	}
	return nil
}

// ToEntityUser transform ReqUpdateProfileByUsername to entity.User.
func (r ReqUpdateProfileByUsername) ToEntityUser() entity.User {
	return entity.User{
		Username: r.Username,
		Password: r.Password,
	}
}
