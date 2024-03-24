package http

import "github.com/Hidayathamir/go-user/pkg/gouser"

// ResUpdatePofile -.
type ResUpdatePofile struct {
	Data  string `json:"data"`
	Error any    `json:"error"`
}

// ResGetProfileByUsername -.
type ResGetProfileByUsername struct {
	Data  gouser.ResGetProfileByUsername `json:"data"`
	Error any                            `json:"error"`
}
