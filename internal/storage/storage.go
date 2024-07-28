package storage

import "errors"

var (
	ErrUserExist    = errors.New("user already exist")
	ErrUserNotFound = errors.New("user not found")
	ErrAppNotfound  = errors.New("app not found")
)
