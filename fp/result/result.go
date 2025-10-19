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

// Success creates a operation from the provided data.
func Success[T any](data T) *Operation[T] {
	return &Operation[T]{
		data: data,
	}
}

// Failure creates an operation from the provided error.
func Failure[T any](err error) *Operation[T] {
	return &Operation[T]{
		err: err,
	}
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

// Map allows you to map on a successful operation to chang the value
// of the "container", will do nothing if op is failed.
func Map[T any, Z any](op *Operation[T], fun func(i T) Z) *Operation[Z] {
	if op.err == nil {
		return &Operation[Z]{fun(op.data), nil}
	}

	var z Z

	return &Operation[Z]{z, op.err}
}

// FlatMap lets you execute a function that already returns a result, with the value of
// of an result. If op is failed, then nothing happens.
func FlatMap[T any, Z any](op *Operation[T], fun func(T) *Operation[Z]) *Operation[Z] {
	if op.err != nil {
		var z Z
		return &Operation[Z]{z, op.err}
	}

	return fun(op.data)
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
