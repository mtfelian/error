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
	FCode    uint   `json:"code"`
	FMessage *string `json:"error,omitempty"`
}

// IStandardError is an interface for standard error
type IStandardError interface {
	error
	Code() uint
	Message() string
}

// Error implements builtin error interface
func (err StandardError) Error() string {
	return fmt.Sprintf("%d: %s", err.Code(), err.Message())
}

// Successful return succes as standard error
func Successful() IStandardError {
	return nil
}

// NewError returns new standard error with code and message from builtin error
func NewError(code uint, err error) IStandardError {
	return StandardError{code, PString(err.Error())}
}

// NewErrorf return new standard error with code, message msg and optional printf args
func NewErrorf(code uint, msg string, args ...interface{}) IStandardError {
	return StandardError{code, PString(fmt.Sprintf(msg, args...))}
}

// MayError makes StandardError from builtin error
func MayError(code uint, err error) IStandardError {
	if err == nil {
		return nil
	}
	return NewError(code, err)
}

// Code returns an error code
func (err StandardError) Code() uint {
	return err.FCode
}

// Message returns an error message
func (err StandardError) Message() string {
	if err.FMessage == nil {
		return ""
	}
	return *err.FMessage
}