package auth

import "errors"

var (
	ErrPermissionDeniedForService  = errors.New("permission denied for service")
	ErrUnableToRetrieveServiceInfo = errors.New("unable to retrieve service info")
	ErrUnableToRetrieveAuth        = errors.New("unable to retrieve auth data")
	ErrInvalidToken                = errors.New("invalid token provided")
	ErrInvalidCredentials          = errors.New("invalid credentials provided")
	ErrUnexpectedSignMethod        = errors.New("unexpected signing method: ")
	ErrInvalidTokenFormat          = errors.New("invalid token format")
	ErrUnableToDecodeToken         = errors.New("unable to decode token")
	ErrUnableToParsePayload        = errors.New("unable to parse payload")
)
