package http

import "github.com/Hidayathamir/go-user/internal/dto"

// contains helper for unit test and integration test.

type resUpdatePofileSuccess struct {
	Data  string `json:"data"`
	Error any    `json:"error"`
}

type resGetProfileByUsernameSuccess struct {
	Data  dto.ResGetProfileByUsername `json:"data"`
	Error any                         `json:"error"`
}
