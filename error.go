package error

import (
	"fmt"
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