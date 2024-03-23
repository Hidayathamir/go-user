package http

import "github.com/Hidayathamir/go-user/internal/dto"

// ResUpdatePofile -.
type ResUpdatePofile struct {
	Data  string `json:"data"`
	Error any    `json:"error"`
}

// ResGetProfileByUsername -.
type ResGetProfileByUsername struct {
	Data  dto.ResGetProfileByUsername `json:"data"`
	Error any                         `json:"error"`
}
