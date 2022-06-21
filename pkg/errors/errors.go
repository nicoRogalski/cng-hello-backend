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
	return errToString(eis.Code, eis.Message)
}

type ErrNotFound struct {
	Code    int
	Message string
}

func NewErrorNotFound(msg string) ErrNotFound {
	return ErrNotFound{Code: 404, Message: msg}
}

func (enf ErrNotFound) Error() string {
	return errToString(enf.Code, enf.Message)
}

type errBadRequest struct {
	Code    int
	Message string
}

func NewErrBadRequest(msg string) errBadRequest {
	return errBadRequest{Code: 400, Message: msg}
}

func (ebr errBadRequest) Error() string {
	return errToString(ebr.Code, ebr.Message)
}

func errToString(code int, message string) string {
	return fmt.Sprintf("Error code: %d, message: %s", code, message)
}
