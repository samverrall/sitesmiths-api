package internal

import (
	"errors"
	"fmt"
)

var (
	ErrBadRequest = errors.New("bad request")
	ErrInternal   = errors.New("internal error")
	ErrNotFound   = errors.New("not found")
)

func WrapErr(serviceErr, causeErr error) error {
	return fmt.Errorf("%w: %s", serviceErr, causeErr)
}
