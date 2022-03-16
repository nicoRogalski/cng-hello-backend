package errors

import (
	"fmt"
)

type ErrInternalServer struct {
	Code    int
	Message string
}

func NewErrInternalServer(msg string) ErrInternalServer {
	return ErrInternalServer{Code: 500, Message: msg}
}

func (eis ErrInternalServer) Error() string {
	return fmt.Sprintf("Error code: %d, message: %s", eis.Code, eis.Message)
}

type ErrNotFound struct {
	Code    int
	Message string
}

func NewErrorNotFound(msg string) ErrNotFound {
	return ErrNotFound{Code: 404, Message: msg}
}

func (enf ErrNotFound) Error() string {
	return fmt.Sprintf("Error code: %d, message: %s", enf.Code, enf.Message)
}
