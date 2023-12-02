package result

import (
	"github.com/Meduzz/helper/fp"
)

type (
	// Operation is inspired by the js promise objects.
	Operation[T any] struct {
		data T
		err  error
	}
)

// Execute take the standard output from any function containerize it.
func Execute[T any](data T, err error) *Operation[T] {
	return &Operation[T]{data, err}
}

// Then execute a function that accepts T and returns (Z, error).
// Will return previous error if present in op.
func Then[T any, Z any](op *Operation[T], fun func(it T) (Z, error)) *Operation[Z] {
	var zero Z

	if op.err != nil {
		return &Operation[Z]{zero, op.err}
	}

	return Execute(fun(op.data))
}

// Recover from an error, by turning it into a T.
// This function is safe to execute even if there's no errors.
func Recover[T any](o *Operation[T], fun func(error) T) *Operation[T] {
	if o.err == nil {
		return o
	}

	return &Operation[T]{fun(o.err), nil}
}

// GetOrElse get T out of o, or run the provided function and return its value.
func GetOrElse[T any](o *Operation[T], fun fp.Producer[T]) T {
	if o.err == nil {
		return o.data
	}

	return fun()
}

// Transform transform an error to another error
func Transform[T any](o *Operation[T], op func(error) error) *Operation[T] {
	if o.err == nil {
		return o
	}

	var zero T

	return &Operation[T]{zero, op(o.err)}
}

// Get return what ever is in o.
func (o *Operation[T]) Get() (T, error) {
	return o.data, o.err
}

// IsSuccess returns true if this operation was a success
func (o *Operation[T]) IsSuccess() bool {
	return o.err == nil
}

// IsFailure returns true if this operation was a failure
func (o *Operation[T]) IsFailure() bool {
	return o.err != nil
}
