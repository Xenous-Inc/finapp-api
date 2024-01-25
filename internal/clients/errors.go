package clients

import "errors"

var (
	ErrRequest         = errors.New("Request error")
	ErrInvalidEntity   = errors.New("Invalid Entity")
	ErrUnauthorized    = errors.New("Unauthorized")
	ErrValidation      = errors.New("Validation")
	ErrNotFound        = errors.New("Entity not found")
	ErrInvalidResponse = errors.New("Invalid response")
	ErrInternal        = errors.New("Internal server error")
)
