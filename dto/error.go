package dto

// Errors represents a stack of error messages
type Errors []interface{}

// Error represents an error that occurred while handling a request
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Errors  Errors `json:"errors,omitempty"`
}

func (e *Error) Error() string {
	return e.Message
}

// NewError generates a new Error instance with an optional message
func NewError(code int, msg ...string) *Error {
	e := &Error{
		Code: code,
	}

	if len(msg) == 0 {
		e.Message = DefaultStatusMessage(code)
		return e
	}

	e.Message = msg[0]
	return e
}

// NewErrorWithStack generates a new Error instance with an error stack and optional message
func NewErrorWithStack(code int, stack Errors, msg ...string) *Error {
	e := NewError(code, msg...)

	e.Errors = stack

	return e
}
