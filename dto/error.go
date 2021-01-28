package dto

import (
	"github.com/dpjungmin/jellypi-server/tools"
)

// Error represents an error that occurred while handling a request.
type Error struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

// ErrorResponse is the formated error response
type ErrorResponse struct {
	Error Error `json:"error"`
}

func (e *Error) Error() string {
	return e.Message
}

// NewError creates a new Error instance with an optional message
func NewError(code int, msg ...string) *Error {
	e := &Error{
		Code: code,
	}
	if len(msg) > 0 {
		e.Message = msg[0]
	} else {
		e.Message = tools.StatusMessage(code)
	}
	return e
}

// NewErrorResponse creates a new ErrorResponse instance with an optional message
func NewErrorResponse(code int, msg ...string) *ErrorResponse {
	return &ErrorResponse{
		Error: *NewError(code, msg...),
	}
}
