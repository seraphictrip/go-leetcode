package lru

import "container/list"

type Node struct {
	key int
	val int
}

type LRUCache struct {
	cap  int
	data map[int]*Node
	// list of nodes, each time use a node
	// move it to back of list
	// so front represents LRU
	list *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:  capacity,
		data: make(map[int]*Node, capacity),
		list: list.New(),
	}
}

func (cache *LRUCache) Get(key int) int {
	node, ok := cache.data[key]
	if !ok {
		return -1
	}
	// TODO: move node to back of list
	return node.val
}

func (cache *LRUCache) Put(key int, value int) {
	if cache.cap == len(cache.data) {
		lru := cache.list.Front()
		node := cache.list.Remove(lru).(*Node)
		delete(cache.data, node.val)
	}
	cache.data[key] = &Node{key, value}
}
