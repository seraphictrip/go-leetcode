package functional

import "slices"

// Creates a new slice equal in length to inputs
// output is populated with the results of calling a provided function on every element of the input array.
//
//	Map(int{1,2,3}, func(n int) int { return n*n})
func Map[T any, V any](inputs []T, transform func(input T) V) []V {
	outputs := make([]V, len(inputs))
	for i := range inputs {
		outputs[i] = transform(inputs[i])
	}
	return outputs
}

func Reduce[T any, V any](inputs []T, reducer func(acc V, input T) V, initialValue V) (output V) {
	output = initialValue
	for i := range inputs {
		output = reducer(output, inputs[i])
	}
	return output
}

// The typical use case is to execute side effects at the end of a chain.
func ForEach[T any](inputs []T, iter func(T)) {
	for i := range inputs {
		iter(inputs[i])
	}
}

func Filter[T any](inputs []T, choice Choose[T]) []T {
	result := make([]T, 0, len(inputs))
	for i := range inputs {
		if choice(inputs[i]) {
			result = append(result, inputs[i])
		}
	}
	return slices.Clip(result)
}

// Check that every item in inputs meets criteria put forth in predicate
// will short circuit if an item does not
func Every[T any](inputs []T, predicate Predicate[T]) bool {
	for i := range inputs {
		if !predicate(inputs[i]) {
			return false
		}
	}
	return true
}

// Check that at least some of the inputs meet criteria in predicat
// will short circuit if true
func Some[T any](inputs []T, predicate Predicate[T]) bool {
	for i := range inputs {
		if predicate(inputs[i]) {
			return true
		}
	}
	return false
}

// Partition of A results in Choice(A) to left of source, and it complement A' to the right
// this can be thought of as putting all choices in one pile (left) and rest in other pile (right)
// Order is not garunteed to be stable
// Partition([1,2,3,4,5], func(n int) bool { return n % 2 == 0}) => [2,4,1,3,5]
func Partition[T any](a []T, choice func(opt T) bool) []T {
	slices.SortFunc(a, func(a, b T) int {
		an, bn := 0, 0
		if choice(a) {
			an = -1
		}
		if choice(b) {
			an = 1
		}
		return an - bn
	})
	return a
}

func PartitionStable[T any](a []T, choice func(opt T) bool) []T {
	slices.SortStableFunc(a, func(a, b T) int {
		an, bn := 0, 0
		if choice(a) {
			an = -1
		}
		if choice(b) {
			an = 1
		}
		return an - bn
	})
	return a
}

// Choice function for set inclusion
// Filter Function if used with Filter
// Paritition Function if used with Partition
// etc...
type Choose[T any] func(T) bool
type ChoiceFn[T any] Choose[T]
type Predicate[T any] Choose[T]

// Transform function, or Mapper
// Transform(T) => V
// T and V can be of the same type
type Transform[T any, V any] func(T) V

// Transforms
func Identity[T any](input T) T {
	return input
}
