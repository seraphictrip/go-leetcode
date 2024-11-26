package arr_test

import (
	"dsa/binarytree/arr"
	"slices"
	"strconv"
	"testing"
)

var InsertTests = []struct {
	inserts []int
}{
	// {},
	// {[]int{1}},
	// {[]int{1, 2}},
	{[]int{1, 2, 3}},
}

func TestInsert(t *testing.T) {
	for i, e := range InsertTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			tree := &arr.BinaryTree[int]{}
			for _, num := range e.inserts {
				tree.Insert(num)
			}

			if !slices.Equal(tree.ToArray(), e.inserts) {
				t.Fatalf("%v != %v", tree.ToArray(), e.inserts)
			}
		})
	}
}

var NOddTests = []struct {
	n   int
	nsq int
}{
	// 1
	{1, 1},
	// 1 + 3
	{2, 4},
	// 1 + 3 + 5
	{3, 9},
	// 1 + 3 + 5 + 7 + 9 + 11 + 13 + 15 + 17
	{9, 81},
	{16, 256},
}

// sum of the first n odd integers is n2
func TestNOdd(t *testing.T) {
	for i, e := range NOddTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if Sum(Seq(1, e.n, 2)) != e.nsq {
				t.Fatalf("%v", e.n)
			}
		})
	}
}

func Seq(start, count, inc int) []int {
	result := make([]int, 0)
	for i := 0; i < count; i++ {
		result = append(result, start+(i*inc))
	}
	return result
}

func Sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
