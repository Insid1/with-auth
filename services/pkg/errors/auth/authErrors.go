package authErrors

import "errors"

var (
	ErrInvalidToken       = errors.New("invalid token provided")
	ErrInvalidCredentials = errors.New("invalid credentials provided")
)
