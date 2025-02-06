package models

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrDuplicateEmail = errors.New("email already exists")
