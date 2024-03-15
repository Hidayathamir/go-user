package http

import "github.com/Hidayathamir/go-user/internal/dto"

// contains helper for unit test and integration test.

type resError struct {
	Data  any    `json:"data"`
	Error string `json:"error"`
}

type resRegisterUserSuccess struct {
	Data  int64 `json:"data"` // userID
	Error any   `json:"error"`
}

type resLoginUserSuccess struct {
	Data  string `json:"data"` // userJWT
	Error any    `json:"error"`
}

type resUpdatePofileSuccess struct {
	Data  string `json:"data"`
	Error any    `json:"error"`
}

type resGetProfileByUsernameSuccess struct {
	Data  dto.ResGetProfileByUsername `json:"data"`
	Error any                         `json:"error"`
}
