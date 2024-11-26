package numislands_test

import (
	"dsa/matrix/numislands"
	"math"
	"slices"
	"strconv"
	"testing"
)

var NumIslandsTests = []struct {
	grid     [][]byte
	expected int
}{
	{},
	{
		[][]byte{
			{'1', '1', '0', '0'},
		},
		1,
	},
	{
		[][]byte{
			{'1', '1', '0', '0'},
			{'1', '1', '0', '0'},
			{'1', '0', '0', '0'},
		},
		1,
	},
	{
		[][]byte{
			{'1', '1', '0', '0'},
			{'1', '1', '0', '0'},
			{'1', '0', '1', '1'},
		},
		2,
	},
	{
		[][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		},
		1,
	},
}

func TestNumIslands(t *testing.T) {
	for i, e := range NumIslandsTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := numislands.NumIslands(e.grid)
			if actual != e.expected {
				t.Fatalf("NumIslands(%v) = %v, want %v", e.grid, actual, e.expected)
			}
		})
	}
}

// x^x = 2^2048

func factors(n int) []int {
	acc := make([]int, 0)
	for n > 1 {
		for i := 2; i < n+1; i++ {
			if n%i == 0 {
				n /= i
				acc = append(acc, i)
				break
			}
		}
	}
	return acc
}

func solve(base, exponent int) int {
	// math.Pow(x, x) = math.Pow(base, exponent)
	x := base
	i := 0
	fs := factors(exponent)
	for math.Pow(float64(x), float64(x)) != math.Pow(float64(base), float64(exponent)) {
		x *= fs[i]
		i++
	}
	return x
}

var SovleTests = []struct {
	base, exponent, expected int
}{
	{2, 2048, 256},
}

func TestSovle(t *testing.T) {
	for i, e := range SovleTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := solve(e.base, e.exponent)
			if actual != e.expected {
				t.Fatalf("solve(%v, %v) = %v, want %v", e.base, e.exponent, actual, e.expected)
			}
		})
	}
}

var factorsTests = []struct {
	n        int
	expected []int
}{
	{},
	{2, []int{2}},
	{3, []int{3}},
	{4, []int{2, 2}},
	{9, []int{3, 3}},
	{2048, []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}},
}

func TestFactors(t *testing.T) {
	for i, e := range factorsTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := factors(e.n)
			if !slices.Equal(actual, e.expected) {
				t.Fatalf("factors(%v) = %v, want %v", e.n, actual, e.expected)
			}
		})
	}
}
