package errors

import (
	"testing"
)

func TestErrorMessage(t *testing.T) {
	err := NewError("test %s", "error")

	if !Is(err, ErrGoSucksBase) {
		t.Errorf("err (%v) is not ErrGoSucksBase", err)
	}

	if !err.Is(ErrGoSucksBase) {
		t.Errorf("err (%v) is not ErrGoSucksBase", err)
	}

	target := NewError("another error")
	if !As(err, &target) {
		t.Errorf("err (%v) is GenericError", err)
	}

	if !err.As(&target) {
		t.Errorf("err (%v) is GenericError", err)
	}

	if err.Error() != "test error" {
		t.Errorf("incorrect error message: %s", err.Error())
	}

	if Unwrap(err) != ErrGoSucksBase {
		t.Errorf("incorrect base error: %v", Unwrap(err))
	}
}

func TestDeriveError(t *testing.T) {
	err := ErrGoSucksBase.Derive("new %s", "error")
	if err.Error() != "new error" {
		t.Errorf("incorrect error message: %s", err.Error())
	}

	if !Is(err, ErrGoSucksBase) {
		t.Errorf("err (%v) is not ErrGoSucksBase", err)
	}
}
