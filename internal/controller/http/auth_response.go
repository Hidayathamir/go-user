package http

import (
	"github.com/Hidayathamir/go-user/pkg/gouser"
)

// ResRegisterUser -.
type ResRegisterUser struct {
	Data  gouser.ResRegisterUser `json:"data"`
	Error any                    `json:"error"`
}

// ResLoginUser -.
type ResLoginUser struct {
	Data  gouser.ResLoginUser `json:"data"`
	Error any                 `json:"error"`
}
