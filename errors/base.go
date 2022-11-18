package errors

type Error interface {
	Error() string
	KindOf(Error) bool
	New(string, ...interface{}) Error
}

type BaseError struct {
	Err  error
	Base string
}
