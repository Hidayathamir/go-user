package http

import "github.com/Hidayathamir/go-user/internal/dto"

// ResRegisterUser -.
type ResRegisterUser struct {
	Data  dto.ResRegisterUser `json:"data"`
	Error any                 `json:"error"`
}
