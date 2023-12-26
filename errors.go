package errors

import "errors"

var (
	ErrUnsupported = errors.ErrUnsupported
	Is             = errors.Is
	Join           = errors.Join
	New            = errors.New
	Unwrap         = errors.Unwrap
)

// As works just like `errors.As`, but it takes care of boilerplate for you.
//
// Instead of:
//
//	var myErr *someError
//	if errors.As(err, &myErr) {
//	   ...
//	}
//
// you can write this:
//
//	if err, ok := errors.As[*someError](err); ok {
//	  ...
//	}
func As[T any](err error) (T, bool) {
	var r T
	return r, errors.As(err, &r)
}
