package fp

import (
	"errors"
)

type (
	// Operation is inspired by the js promise objects.
	Operation[T comparable] struct {
		data T
		err  error
	}
)

var (
	ErrConvert = errors.New("failed to convert variable")
)

// Execute take the standard output from any function containerize it.
func Execute[T comparable](data T, err error) *Operation[T] {
	return &Operation[T]{data, err}
}

// Then execute a function that accepts T and returns (Z, error).
// Will return previous error if present in op.
func Then[T comparable, Z comparable](op *Operation[T], fun func(it T) (Z, error)) *Operation[Z] {
	if op.err != nil {
		return &Operation[Z]{*new(Z), op.err}
	}

	return Execute(fun(op.data))
}

// Filter run a filter on the value of o, that returns o if it passed or returns an empty Operation[T] otherwise.
func (o *Operation[T]) Filter(fun func(T) bool) *Operation[T] {
	if fun(o.data) {
		return o
	}

	return &Operation[T]{*new(T), nil}
}

// Recover from an error, by turning it into a T.
// This function is safe to execute even if there's no errors.
func (o *Operation[T]) Recover(fun func(error) T) *Operation[T] {
	if o.err == nil {
		return o
	}

	return &Operation[T]{fun(o.err), nil}
}

// GetOrElse get T out of o, or run the provided function and return its value.
func (o *Operation[T]) GetOrElse(fun func() T) T {
	var dead T
	if o.data != dead {
		return o.data
	}

	return fun()
}

// Get return what ever is in o.
func (o *Operation[T]) Get() (T, error) {
	return o.data, o.err
}
