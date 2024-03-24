package http

import "github.com/Hidayathamir/go-user/internal/usecase"

// ResUpdatePofile -.
type ResUpdatePofile struct {
	Data  string `json:"data"`
	Error any    `json:"error"`
}

// ResGetProfileByUsername -.
type ResGetProfileByUsername struct {
	Data  usecase.ResGetProfileByUsername `json:"data"`
	Error any                             `json:"error"`
}
