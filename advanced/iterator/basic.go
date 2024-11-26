package iterator

import "iter"

func Until(n int) []int {
	result := make([]int, 0, n)
	for i := range n {
		result = append(result, i)
	}
	return result
}

func LazyCount(n int) func(func(int) bool) {
	return func(yield func(i int) bool) {
		for i := range n {
			if !yield(i) {
				return
			}
		}

	}
}

type Slice[T any, V any] []T

func (s Slice[T, V]) Map(transform func(item T) V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range s {
			if !yield(transform(v)) {
				return
			}
		}
	}
}
