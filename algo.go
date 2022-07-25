package algos

type mapIterator[T any] struct {
	source Iterator[T]
	mapper func(T) T
}

// Get next element
func (iter *mapIterator[T]) Next() bool {
	return iter.source.Next()
}

// Get value of current elemented visisted by iterator
func (iter *mapIterator[T]) Value() T {
	value := iter.source.Value()
	return iter.mapper(value)
}

func Map[T any](iter Iterator[T], f func(T) T) Iterator[T] {
	return &mapIterator[T]{
		iter, f,
	}
}

type filterIterator[T any] struct {
	source Iterator[T]
	pred   func(T) bool
}

// Get next element
func (iter *filterIterator[T]) Next() bool {
	for iter.source.Next() {
		if iter.pred(iter.source.Value()) {
			return true
		}
	}
	return false
}

// Get value of current elemented visisted by iterator
func (iter *filterIterator[T]) Value() T {
	return iter.source.Value()
}

func Filter[T any](iter Iterator[T], pred func(T) bool) Iterator[T] {
	return &filterIterator[T]{
		iter, pred,
	}
}

func Collect[T any](iter Iterator[T]) []T {
	var xs []T

	for iter.Next() {
		xs = append(xs, iter.Value())
	}

	return xs
}

// Function provided to Reduce higher order function
// values get combined with accumulator and a new accumulated
// value gets returned
type Reducer[T, V any] func(accum T, value V) T

// Reduce values iterated over to a single value
func Reduce[T, V any](iter Iterator[V], f Reducer[T, V]) T {
	var accum T
	for iter.Next() {
		accum = f(accum, iter.Value())
	}
	return accum
}
