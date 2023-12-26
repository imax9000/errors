// Package errors is a wrapper for `errors` package from standard library
// that takes advantage of recently introduced generics.
//
// Function As is replaced to allow the following:
//
//	if err, ok := errors.As[*someError](err); ok {
//	  ...
//	}
//
// Rest of the standard `errors` package is re-exported as is.
package errors
