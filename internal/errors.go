package internal

import "errors"

var (
	ErrBadRequest = errors.New("bad request")
	ErrInternal   = errors.New("internal error")
	ErrNotFound   = errors.New("not found")
)
