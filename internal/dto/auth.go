package dto

import (
	"errors"

	"github.com/Hidayathamir/go-user/internal/entity"
)

// ReqLoginUser -.
type ReqLoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate validate ReqLoginUser.
func (r ReqLoginUser) Validate() error {
	if r.Username == "" {
		return errors.New("ReqLoginUser.Username can not be empty")
	}
	if r.Password == "" {
		return errors.New("ReqLoginUser.Password can not be empty")
	}
	return nil
}

// ReqRegisterUser -.
type ReqRegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate validate ReqRegisterUser.
func (r ReqRegisterUser) Validate() error {
	if r.Username == "" {
		return errors.New("ReqRegisterUser.Username can not be empty")
	}
	if r.Password == "" {
		return errors.New("ReqRegisterUser.Password can not be empty")
	}
	return nil
}

// ToEntityUser transform ReqRegisterUser to entity.User.
func (r ReqRegisterUser) ToEntityUser() entity.User {
	return entity.User{
		Username: r.Username,
		Password: r.Password,
	}
}

// ResRegisterUser -.
type ResRegisterUser struct {
	UserID int64 `json:"user_id"`
}
