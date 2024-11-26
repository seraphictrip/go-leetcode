package core

type DynamicArray[T any] []T

func (a DynamicArray[T]) Get(i int) T {
	return a[i]
}

func (a DynamicArray[T]) Set(i int, val T) {
	a[i] = val
}

func (a *DynamicArray[T]) PushBack(val T) {
	if cap(*a) == len(*a) {

	}
}
