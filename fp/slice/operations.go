package slice

import (
	"reflect"

	"github.com/Meduzz/helper/fp"
)

// Map on a slice of T returning a slice of K
func Map[T any, K any](in []T, handler func(T) K) []K {
	out := make([]K, 0)

	for _, i := range in {
		out = append(out, handler(i))
	}

	return out
}

// FlatMap on a slice of T to a slice of K
func FlatMap[T any, K any](in []T, handler func(T) []K) []K {
	out := make([]K, 0)

	for _, i := range in {
		out = append(out, handler(i)...)
	}

	return out
}

// Fold over a slice of T merging into type of K
func Fold[T any, K any](in []T, agg K, handler func(T, K) K) K {
	for _, i := range in {
		agg = handler(i, agg)
	}

	return agg
}

// ForEach over a slice of T
func ForEach[T any](in []T, handler fp.Consumer[T]) {
	for _, i := range in {
		handler(i)
	}
}

// Filter over a slice of T returning all matches
func Filter[T any](in []T, handler fp.Predicate[T]) []T {
	type K []T
	out := make([]T, 0)

	out = Fold[T, K](in, out, func(t T, k K) K {
		if handler(t) {
			return append(k, t)
		}

		return k
	})

	return out
}

// Concat concats 2 slices of T into one
func Concat[T any](first []T, second []T) []T {
	return append(first, second...)
}

// Head returns first value of slice of T
func Head[T any](in []T) T {
	if len(in) == 0 {
		var out T
		return out
	}

	return in[0]
}

// Tail returns all but first item of slice of T
func Tail[T any](in []T) []T {
	if len(in) == 0 {
		return make([]T, 0)
	}

	return in[1:]
}

// Take returns count from slice of T or all of T is less than count
func Take[T any](in []T, count int) []T {
	if len(in) == 0 {
		return make([]T, 0)
	} else if len(in) < count {
		return in
	}

	return in[:count]
}

// Skip returns everything from slice of T after count
func Skip[T any](in []T, count int) []T {
	if len(in) == 0 {
		return make([]T, 0)
	} else if len(in) < count {
		return in
	}

	return in[count:]
}

// Partition splits slice of T into chunks of size
func Partition[T any](in []T, size int) [][]T {
	out := make([][]T, 0)
	step := make([]T, 0)

	ForEach(in, func(t T) {
		step = append(step, t)

		if len(step) == size {
			out = append(out, step)
			step = make([]T, 0)
		}
	})

	if len(step) != 0 {
		return append(out, step)
	} else {
		return out
	}
}

// Group a slice of T into a map[string][]T
func Group[T any](in []T, grouper func(T) string) map[string][]T {
	out := make(map[string][]T)

	Fold(in, out, func(t T, m map[string][]T) map[string][]T {
		key := grouper(t)
		items := m[key]

		if items == nil {
			items = make([]T, 0)
		}

		items = append(items, t)
		m[key] = items

		return m
	})

	return out
}

// Shard a slice of T into a map[int64][]T
func Shard[T any](in []T, sharder func(T) int64) map[int64][]T {
	out := make(map[int64][]T, 0)

	Fold(in, out, func(t T, m map[int64][]T) map[int64][]T {
		key := sharder(t)
		list := m[key]

		if list == nil {
			list = make([]T, 0)
		}

		list = append(list, t)

		m[key] = list

		return m
	})

	return out
}

// Contains over a slice of T returning true if needle exists in slice
func Contains[T any](in []T, needle T) bool {
	return 0 < len(Filter(in, func(t T) bool {
		return reflect.DeepEqual(t, needle)
	}))
}
