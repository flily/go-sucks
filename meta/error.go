package meta

import (
	"reflect"

	"github.com/flily/go-sucks/errors"
)

var (
	ErrMetaError      = errors.NewError("meta error")
	ErrNotDuplcatable = ErrMetaError.Derive("not duplicatable")
	ErrUntypedNil     = ErrMetaError.Derive("untyped nil is unacceptable")
)

func NewMetaError(format string, args ...interface{}) error {
	return ErrMetaError.Derive(format, args...)
}

func NewNotDuplicatableError(kind reflect.Kind) error {
	return ErrNotDuplcatable.Derive("kind %s is not duplicatable", kind)
}
