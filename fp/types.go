package fp

type (
	Producer[T any]           func() T
	Consumer[T any]           func(T)
	Predicate[T any]          func(T) bool
	Operation[T any, K any]   func(T) (K, error)
	Action[T any]             func(T) error
	Calculation[T any, K any] func(T) K
)

// And - combine 2 predicates into one with &&
func And[T any](first, second Predicate[T]) Predicate[T] {
	return func(t T) bool {
		return first(t) && second(t)
	}
}

// Or - combine 2 predicates into one with ||
func Or[T any](first, second Predicate[T]) Predicate[T] {
	return func(t T) bool {
		return first(t) || second(t)
	}
}
