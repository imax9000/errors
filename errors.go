package errors

import "errors"

var (
	ErrUnsupported = errors.ErrUnsupported
	Is             = errors.Is
	Join           = errors.Join
	New            = errors.New
	Unwrap         = errors.Unwrap
)

func As[T any](err error) (T, bool) {
	var r T
	return r, errors.As(err, &r)
}
