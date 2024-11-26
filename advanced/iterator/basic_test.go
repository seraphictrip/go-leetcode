package iterator_test

import (
	"dsa/advanced/iterator"
	"fmt"
	"slices"
	"strconv"
	"testing"
)

var UntilTests = []struct {
	n        int
	expected []int
}{
	{100, nil},
}

func TestUntil(t *testing.T) {
	for i, e := range UntilTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := iterator.Until(e.n)
			if !slices.Equal(actual, e.expected) {
				t.Fatalf("Until(%v) = %v, want %v", e.n, actual, e.expected)
			}
		})
	}
}

var LazyCountTests = []struct {
	n int
}{
	{100},
}

func TestLazyCount(t *testing.T) {
	for i, e := range LazyCountTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			last := 0
			for i := range iterator.LazyCount(e.n) {
				last = i
			}
			if last != e.n-1 {
				t.Fatalf("LazyCount(%v) counts to %v, want %v", e.n, last, e.n-1)
			}
		})
	}
}

var MapTests = []struct {
	arr      []int
	expected []rune
}{
	{
		[]int{int('a'), 33, 3, 4, 5}, nil,
	},
}

func TestMap(t *testing.T) {
	for i, e := range MapTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			s := iterator.Slice[int, string](e.arr)
			for val := range s.Map(func(num int) string {
				return string(rune(num))
			}) {
				fmt.Println(val)
			}
		})
	}
}
