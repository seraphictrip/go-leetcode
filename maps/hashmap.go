package maps

import (
	"errors"
)

type Map[T comparable, V any] interface {
	Put(key T, val V)
	Get(key T) (V, bool)
}

type Basic[T comparable, V any] map[T]V

func (b Basic[T, V]) Put(key T, val V) {
	b[key] = val
}

func (b Basic[T, V]) Get(key T) (V, bool) {
	val, ok := b[key]
	return val, ok
}

type Pair[Val any] struct {
	key string
	val Val
}

func NewPair[Val any](key string, val Val) *Pair[Val] {
	return &Pair[Val]{
		key: key,
		val: val,
	}
}

type HashMap[Val any] struct {
	size int
	cap  int
	data []*Pair[Val]
}

func NewHashMap[Val any](cap int) *HashMap[Val] {
	return &HashMap[Val]{
		// number of actual entries
		size: 0,
		// current capacity
		cap: cap,
		// KeyValue pairs
		data: make([]*Pair[Val], cap),
	}
}

func (m HashMap[Val]) Put(key string, val Val) {
	index := m.hash(key)
	for {
		if m.data[index] == nil {
			// no collision, add item and expand if necessary
			m.data[index] = NewPair(key, val)
			m.size += 1
			if m.size >= m.cap/2 {
				m.rehash()
			}
			return
		} else if m.data[index].key == key {
			m.data[index].val = val
			return
		}
		// look for close
		index += 1
		index = index % m.cap
	}
}

func (m HashMap[Val]) Get(key string) (v Val, err error) {
	index := m.hash(key)
	for m.data[index] != nil {
		if m.data[index].key == key {
			return m.data[index].val, nil
		}
		index += 1
		index = index % m.cap
	}
	return v, errors.New("not found")
}

func (m HashMap[Val]) hash(key string) int {
	index := 0
	for _, ch := range key {
		index += int(ch)
	}
	return index % m.cap
}

func (m *HashMap[Val]) rehash() {
	m.cap = 2 * m.cap
	newMap := make([]*Pair[Val], m.cap)
	oldMap := m.data
	m.data = newMap
	m.size = 0
	for _, pair := range oldMap {
		if pair != nil {
			m.Put(pair.key, pair.val)
		}
	}
}
