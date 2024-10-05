package user

import "errors"

var (
	ErrUnableToCreate    = errors.New("unable to create user")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrUserAlreadyExist  = errors.New("user already exist")
	ErrUserNotFound      = errors.New("user not found")
	ErrUserIDNotProvided = errors.New("id is not provided")
	ErrEmptyPassword     = errors.New("empty password provided")
)
