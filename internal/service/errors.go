package service

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrStop     = errors.New("load stop func error")
)
