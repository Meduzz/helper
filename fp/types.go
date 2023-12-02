package fp

type (
	Producer[T any]  func() T
	Consumer[T any]  func(T)
	Predicate[T any] func(T) bool
)
