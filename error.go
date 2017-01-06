package error

import (
	"fmt"
	. "github.com/mtfelian/utils"
)

// CodeSuccess is a success error code
const (
	CodeSuccess uint = iota
)

// StandardError is a standard error to return with Gin
type StandardError struct {
	Code    uint   `json:"code"`
	Message *string `json:"error,omitempty"`
}

// Error implements builtin error interface
func (err StandardError) Error() string {
	if err.Message != nil {
		return fmt.Sprintf("%d: %s", err.Code, *err.Message)
	}
	return fmt.Sprintf("%d", err.Code)
}

// Occurred return true if it is an error, otherwise returns false,
// this check is analogous to (err != nil)
func (err StandardError) Occurred() bool {
	return err.Code != CodeSuccess
}

// Successful return succes as standard error
func Successful() StandardError {
	return StandardError{CodeSuccess, nil}
}

// NewError returns new standard error with code and message from builtin error
func NewError(code uint, err error) StandardError {
	return StandardError{code, PString(err.Error())}
}

// NewErrorf return new standard error with code, message msg and optional printf args
func NewErrorf(code uint, msg string, args ...interface{}) StandardError {
	return StandardError{code, PString(fmt.Sprintf(msg, args...))}
}

// MayError makes StandardError from builtin error
func MayError(code uint, err error) StandardError {
	if err != nil {
		return NewError(code, err)
	}
	return Successful()
}
