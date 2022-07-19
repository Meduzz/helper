package fp

type (
	// Generix it's a generic array
	Generix[T any] []T
	K              any
)

// From generate a Generix[T] from an array
func From[T any](array []T) Generix[T] {
	return array
}

// Of generate a Generix[T] from an bunch of T:s.
func Of[T any](data ...T) Generix[T] {
	return data
}

// Filter apply a filter function to a Generix[T]
func (g Generix[T]) Filter(filter func(it T) bool) Generix[T] {
	keepers := make([]T, 0)

	for _, val := range g {
		if filter(val) {
			keepers = append(keepers, val)
		}
	}

	return From(keepers)
}

// Map apply a map function to a bunch of T:s
// Deprecated: use fp.Map instead.
func (g Generix[T]) Map(mapper func(it T) K) Generix[K] {
	changed := make([]K, 0)

	for _, val := range g {
		item := mapper(val)
		changed = append(changed, item)
	}

	return From(changed)
}

// FlatMap apply a flatMp function to a Generix[T]
// Deprecated: use fp.FlatMap isntead.
func (g Generix[T]) FlatMap(mapper func(it T) []K) Generix[K] {
	changed := make([]K, 0)

	for _, val := range g {
		items := mapper(val)
		changed = append(changed, items...)
	}

	return From(changed)
}

// Fold turn a Generix[T] into a K
// Deprecated: use fp.Fold instead.
func (g Generix[T]) Fold(target K, fold func(it T, agg K) K) K {
	for _, val := range g {
		target = fold(val, target)
	}

	return target
}

// ForEach apply a function that reads each item in a Generix[T]
func (g Generix[T]) ForEach(fun func(T)) {
	for _, val := range g {
		fun(val)
	}
}

// Concat merge 2 Generix[T] into one.
func (g Generix[T]) Concat(it Generix[T]) Generix[T] {
	return append(g, it...)
}

// Map apply a map function (T)=>K to a Generix[T] to get a Generix[K]
func Map[T any, K any](it Generix[T], fun func(T) K) Generix[K] {
	changed := make([]K, 0)

	for _, val := range it {
		changed = append(changed, fun(val))
	}

	return From(changed)
}

// FlatMap apply a flatMap function (T)=>[]K to a Generix[T] to get a Generix[K]
func FlatMap[T any, K any](it Generix[T], fun func(T) []K) Generix[K] {
	changed := make([]K, 0)

	for _, val := range it {
		ret := fun(val)
		changed = append(changed, ret...)
	}

	return From(changed)
}

// Fold apply a fold function to fold a Generix[T] to a single K.
func Fold[T any, K any](it Generix[T], agg K, fun func(T, K) K) K {
	for _, val := range it {
		agg = fun(val, agg)
	}

	return agg
}

// TODO partition (partitionSize)[][]K
// TODO group (func(it T) string)map[string][]K