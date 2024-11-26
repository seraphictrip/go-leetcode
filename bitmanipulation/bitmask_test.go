package bitmanipulation_test

import (
	bm "dsa/bitmanipulation"
	"fmt"
	"strconv"
	"testing"
)

var ConstantsTests = []struct {
	constant bm.Bits
	expected bm.Bits
}{
	// 0b00000000
	{0, 0},
	{bm.Empty, 0},
	// 0b00000001
	{bm.F0, 1},
	{bm.F0, 0b00000001},
	{bm.F0, 0b1},
	// 0b00000010
	{bm.F1, 2},
	{bm.F1, 0b00000010},
	{bm.F1, 0b10},
	// 0b00000100
	{bm.F2, 4},
	{bm.F2, 0b00000100},
	// 0b00001000
	{bm.F3, 8},
	{bm.F3, 0b00001000},
	// 0b00010000
	{bm.F4, 16},
	{bm.F4, 0b00010000},
	{bm.F4, 0b10000},
	// 0b00100000
	{bm.F5, 32},
	// 0b01000000
	{bm.F6, 64},
	// 0b10000000
	{bm.F7, 128},
	// 0b00000001 | 0b00000010 => 0b00000011
	{bm.F0 | bm.F1, 3},
	{1 | 2, 3},
	// 0b00000010 | 0b00000001 => 0b00000011
	{bm.F1 | bm.F0, 3},
	{2 | 1, 3},
	{bm.Set(bm.F1, bm.F0), 3},
	// full 0b11111111
	{bm.Full, 255},
}

func TestConstants(t *testing.T) {
	for i, e := range ConstantsTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if e.constant != e.expected {
				t.Fatalf("%08b != %08b", e.constant, e.expected)
			}
		})
	}
}

var HasTests = []struct {
	mask, flags bm.Bits
	expected    bool
}{
	// 0b00000000 has nothing
	{0, bm.F0, false},
	{0, bm.F1, false},
	{0, bm.F2, false},
	{0, bm.F3, false},
	{0, bm.F4, false},
	{0, bm.F5, false},
	{0, bm.F6, false},
	{0, bm.F7, false},
	// 0b11111111 has everything
	{255, bm.F0, true},
	{255, bm.F1, true},
	{255, bm.F2, true},
	{255, bm.F3, true},
	{255, bm.F4, true},
	{255, bm.F5, true},
	{255, bm.F6, true},
	{255, bm.F7, true},
	// everthing contains itself
	{bm.F0, bm.F0, true},
	{bm.F1, bm.F1, true},
	{bm.F2, bm.F2, true},
	{bm.F3, bm.F3, true},
	{bm.F4, bm.F4, true},
	{bm.F5, bm.F5, true},
	{bm.F6, bm.F6, true},
	{bm.F7, bm.F7, true},
	// I do not have my complement complement
	{bm.F0, 0b11111110, false},
	{bm.F0, bm.F0 ^ 255, false},
	{bm.F0, bm.Complement(bm.F0), false},
	{bm.F1, 0b11111101, false},
	{bm.F1, bm.F1 ^ 255, false},
	{bm.F2, bm.Complement(bm.F2), false},
	{bm.F3, bm.Complement(bm.F3), false},
	{bm.F3, bm.Complement(bm.F3), false},
	{bm.F4, bm.Complement(bm.F4), false},
	{bm.F5, bm.Complement(bm.F5), false},
	{bm.F6, bm.Complement(bm.F6), false},
	{bm.F7, bm.Complement(bm.F7), false},
	// my complement does not have me
	{bm.Complement(bm.F0), bm.F0, false},
	{bm.Complement(bm.F1), bm.F1, false},
	{bm.Complement(bm.F2), bm.F2, false},
	{bm.Complement(bm.F3), bm.F3, false},
	{bm.Complement(bm.F4), bm.F4, false},
	{bm.Complement(bm.F5), bm.F5, false},
	{bm.Complement(bm.F6), bm.F6, false},
	{bm.Complement(bm.F7), bm.F7, false},
	// NOTE: has is best uesd for single flags, use ContainsAll if looking to match a flag set
}

func TestHas(t *testing.T) {
	for i, e := range HasTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := bm.Has(e.mask, e.flags)
			fmt.Printf("Has(%08b, %08b) = %v, want %v\n", e.mask, e.flags, actual, e.expected)
			if actual != e.expected {
				t.Fatalf("Has(%08b, %08b) = %v, want %v", e.mask, e.flags, actual, e.expected)
			}
		})
	}
}

var SetTests = []struct {
	mask, flag, expected bm.Bits
}{
	// empty and any single flag is flag
	{bm.Empty, bm.F0, bm.F0},
	{bm.Empty, bm.F1, bm.F1},
	{bm.Empty, bm.F2, bm.F2},
	{bm.Empty, bm.F3, bm.F3},
	{bm.Empty, bm.F4, bm.F4},
	{bm.Empty, bm.F5, bm.F5},
	{bm.Empty, bm.F6, bm.F6},
	{bm.Empty, bm.F7, bm.F7},
	// setting self on self doesn't change anything
	{bm.F0, bm.F0, bm.F0},
	{bm.F1, bm.F1, bm.F1},
	{bm.F2, bm.F2, bm.F2},
	{bm.F3, bm.F3, bm.F3},
	{bm.F4, bm.F4, bm.F4},
	{bm.F5, bm.F5, bm.F5},
	{bm.F6, bm.F6, bm.F6},
	{bm.F7, bm.F7, bm.F7},
	// set multiple flags as once
	{bm.Empty, bm.F0 | bm.F1, 0b00000011},
	{bm.F0, bm.F0 | bm.F1, 0b00000011},
}

func TestSet(t *testing.T) {
	for i, e := range SetTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// Set sets 1 or more flags
			actual := bm.Set(e.mask, e.flag)
			if actual != e.expected {
				t.Fatalf("Set(%08b, %08b) = %08b, want %08b", e.mask, e.flag, actual, e.expected)
			}
		})
	}
}

var EqualityTests = []struct {
	actual, expected bm.Bits
}{
	// anything union with complement is full set
	{bm.Set(bm.F0, bm.Complement(bm.F0)), bm.Full},
	{bm.Set(bm.F1, bm.Complement(bm.F1)), bm.Full},
	{bm.Set(bm.F2, bm.Complement(bm.F2)), bm.Full},
	{bm.Set(bm.F3, bm.Complement(bm.F3)), bm.Full},
	{bm.Set(bm.F4, bm.Complement(bm.F4)), bm.Full},
	{bm.Set(bm.F5, bm.Complement(bm.F5)), bm.Full},
	{bm.Set(bm.F6, bm.Complement(bm.F6)), bm.Full},
	{bm.Set(bm.F7, bm.Complement(bm.F7)), bm.Full},
	// anything union with self, is self
	{bm.Set(bm.F0, bm.F0), bm.F0},
	{bm.Set(bm.F1, bm.F1), bm.F1},
	{bm.Set(bm.F2, bm.F2), bm.F2},
	{bm.Set(bm.F3, bm.F3), bm.F3},
	{bm.Set(bm.F4, bm.F4), bm.F4},
	{bm.Set(bm.F5, bm.F5), bm.F5},
	{bm.Set(bm.F6, bm.F6), bm.F6},
	{bm.Set(bm.F7, bm.F7), bm.F7},
	{bm.Complement(bm.Empty), bm.Full},
	{bm.Complement(bm.Full), bm.Empty},
}

func TestEquality(t *testing.T) {
	for i, e := range EqualityTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if e.actual != e.expected {
				t.Fatalf("%08b != %08b", e.actual, e.expected)
			}
		})
	}
}

var ContainsAllTests = []struct {
	mask, flags bm.Bits
	expected    bool
}{
	// empty will not contain
	{bm.Empty, bm.Full, false},
	{bm.Empty, bm.F0, false},
	// full contains everything (even empty set)
	{bm.Full, bm.Empty, true},
	{bm.Full, bm.Full, true},
	{bm.Full, bm.F0, true},
	{bm.Full, bm.F1, true},

	{bm.F0 | bm.F1, 0b00000011, true},
	{bm.F0 | bm.F1 | bm.F2, 0b00000111, true},
	{bm.F0 | bm.F1 | bm.F2 | bm.F3, 0b00001111, true},
	{bm.F0 | bm.F1 | bm.F2 | bm.F3 | bm.F4, 0b00011111, true},
	{bm.F0 | bm.F1 | bm.F2 | bm.F3 | bm.F4 | bm.F5, 0b00111111, true},
	{bm.F0 | bm.F1 | bm.F2 | bm.F3 | bm.F4 | bm.F5 | bm.F6, 0b01111111, true},
	{bm.F0 | bm.F1 | bm.F2 | bm.F3 | bm.F4 | bm.F5 | bm.F6 | bm.F7, 0b11111111, true},

	// must containall
	{0b11111111, bm.Full, true},
	{0b00111111, bm.Full, false},
	{0b11011111, bm.Full, false},
	{0b11101111, bm.Full, false},
}

func TestContainsAll(t *testing.T) {
	for i, e := range ContainsAllTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := bm.ContainsAll(e.mask, e.flags)
			if actual != e.expected {
				t.Fatalf("ContainsAll(%v, %v) = %v, want %v", e.mask, e.flags, actual, e.expected)
			}
		})
	}
}

var countBitsTests = []struct {
	n        uint
	exepcted int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 1},
	{5, 2},
	{6, 2},
	{7, 3},
	{8, 1},
	{9, 2},
	{10, 2},
	{11, 3},
	{12, 2},
	{13, 3},
	{14, 3},
	{15, 4},
	{16, 1},
}

func TestCountBits(t *testing.T) {
	for i, e := range countBitsTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := countBits(e.n)
			if actual != e.exepcted {
				t.Fatalf("countBits(%v) = %v, want %v", e.n, actual, e.exepcted)
			}
		})
	}
}

func countBits(n uint) int {
	count := 0
	for n > 0 {
		if n&1 == 1 {
			count++
		}
		n = n >> 1
	}

	return count
}
