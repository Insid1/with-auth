package userErrors

import "errors"

var (
	ErrUnableToCreate   = errors.New("unable to create user")
	ErrInvalidPassword  = errors.New("invalid password")
	ErrUserAlreadyExist = errors.New("user already exist")
)
