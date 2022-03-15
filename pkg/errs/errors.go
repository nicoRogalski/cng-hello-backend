package errs

import "errors"

var (
	ErrNotFound       = errors.New("ressource not found")
	ErrInternalServer = errors.New("internal server error")
)
