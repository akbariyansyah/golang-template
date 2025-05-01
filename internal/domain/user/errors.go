package user

import "errors"

var (
	ErrEmailNotFound     = errors.New("user: email not found")
	ErrEmailEmpty        = errors.New("user: email cannot be empty")
	ErrEmailAlreadyExist = errors.New("user: email already exist")
)
