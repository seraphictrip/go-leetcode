package maps_test

import (
	"dsa/maps"
	"strconv"
	"testing"
)

type keyvalue struct {
	key string
	val int
}

var InsertBasicTests = []struct {
	m     maps.Basic[string, int]
	pairs []keyvalue
}{
	{
		maps.Basic[string, int](map[string]int{}),
		[]keyvalue{{"test", 1}},
	},
	{
		maps.Basic[string, int](map[string]int{}),
		[]keyvalue{{"test", 1}, {"this", 2}, {"out", 3}},
	},
}

func TestInsertBasic(t *testing.T) {
	for i, e := range InsertBasicTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			for _, kv := range e.pairs {
				e.m.Put(kv.key, kv.val)
				val, ok := e.m.Get(kv.key)
				if !ok || val != kv.val {
					t.Fatalf("%+v missing %v:%v", e.m, kv.key, kv.val)
				}
			}

		})
	}
}
