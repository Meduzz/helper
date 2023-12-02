package result

import "reflect"

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
func GetOrElse[T any](o *Operation[T], fun func() T) T {
	if !reflect.ValueOf(o.data).IsZero() {
		return o.data
	}

	return fun()
}

// Get return what ever is in o.
func (o *Operation[T]) Get() (T, error) {
	return o.data, o.err
}
