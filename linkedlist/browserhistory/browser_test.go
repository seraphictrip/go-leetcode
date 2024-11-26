package browserhistory

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestConstructor(t *testing.T) {
	bh := Constructor("leetcode.com")
	fmt.Println(bh)
}

var VisitTests = []struct {
	visits []string
}{
	{[]string{"facebook.com", "google.come"}},
}

func TestVisit(t *testing.T) {
	for i, e := range VisitTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Constructor("leetcode.com")
			bh := &b
			for i := range e.visits {
				bh.Visit(e.visits[i])
				if bh.Current() != e.visits[i] {
					t.Fatalf("bh.Visit(%v) failed? %v", e.visits[i], bh)
				}
			}

		})
	}
}

var BackTests = []struct {
	initial  []string
	steps    int
	expected string
}{
	{[]string{"a", "b", "c", "d", "e"}, 2, "c"},
	{strings.Split("abcdefghijklmnopqrstuvwxyz", ""), 14, "l"},
}

func TestBack(t *testing.T) {
	for i, e := range BackTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Constructor(e.initial[0])
			bh := &b
			for i := 1; i < len(e.initial); i++ {
				bh.Visit(e.initial[i])
				if bh.Current() != e.initial[i] {
					t.Fatalf("bh.Visit(%v) failed? %v", e.initial[i], bh)
				}
			}
			actual := bh.Back(e.steps)
			if actual != e.expected {
				t.Fatalf("bh.Back(%v) = %v, want %v", e.steps, actual, e.expected)
			}
		})
	}
}
