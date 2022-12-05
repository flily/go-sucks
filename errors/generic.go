package errors

import (
	sys_err "errors"
	"fmt"
)

var (
	// ErrGoSucksBase is the base error for all errors in this library.
	ErrGoSucksBase = &GenericError{
		Base:    nil,
		Message: "base error of library go-sucks",
	}
)

// Is is equivalent to system's errors.Is.
func Is(err error, target error) bool {
	return sys_err.Is(err, target)
}

// As is equivalent to system's errors.As.
func As(err error, target interface{}) bool {
	return sys_err.As(err, target)
}

// Unwrap is equivalent to system's errors.Unwrap.
func Unwrap(err error) error {
	return sys_err.Unwrap(err)
}

type UsableError interface {
	error
	Unwrap() error
	Derive(format string, args ...interface{}) UsableError
	Is(target error) bool
	As(target interface{}) bool
}

// GenericError is the error type for generic use.
type GenericError struct {
	Base    error
	Message string
	Data    interface{}
}

func (e *GenericError) Error() string {
	return e.Message
}

func (e *GenericError) Unwrap() error {
	return e.Base
}

func (e *GenericError) Is(target error) bool {
	return Is(e.Base, target)
}

func (e *GenericError) As(target interface{}) bool {
	return As(e.Base, target)
}

func (e *GenericError) Derive(format string, args ...interface{}) UsableError {
	err := &GenericError{
		Base:    e,
		Message: fmt.Sprintf(format, args...),
	}
	return err
}

func DeriveError(base error, format string, args ...interface{}) UsableError {
	err := &GenericError{
		Base:    base,
		Message: fmt.Sprintf(format, args...),
	}
	return err
}

// NewError creates a new error from ErrGoSucksBase.
func NewError(format string, args ...interface{}) UsableError {
	return DeriveError(ErrGoSucksBase, format, args...)
}
