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
	Message string `json:"error"`
}

// Error implements builtin error interface
func (err StandardError) Error() string {
	return fmt.Sprintf("%d: %s", err.Code, err.Message)
}
