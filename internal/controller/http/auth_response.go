package http

import "github.com/Hidayathamir/go-user/internal/usecase"

// ResRegisterUser -.
type ResRegisterUser struct {
	Data  usecase.ResRegisterUser `json:"data"`
	Error any                     `json:"error"`
}

// ResLoginUser -.
type ResLoginUser struct {
	Data  usecase.ResLoginUser `json:"data"`
	Error any                  `json:"error"`
}
