// Package gouser -.
package gouser

import "errors"

var (
	// ErrRequestInvalid occurs when request invalid.
	ErrRequestInvalid = errors.New("request invalid")
	// ErrJWTAuth occurs when there is a problem with JWT auth.
	ErrJWTAuth = errors.New("JWT auth")
	// ErrNothingToBeUpdate occurs when nothing to be update.
	ErrNothingToBeUpdate = errors.New("nothing to be update")
	// ErrWrongPassword occurs when user login with wrong password.
	ErrWrongPassword = errors.New("wrong password")
	// ErrDuplicateUsername occurs when register user but username already exists.
	ErrDuplicateUsername = errors.New("duplicate username")
	// ErrUnknownUsername occurs when username does not exists.
	ErrUnknownUsername = errors.New("unknown username")
)
