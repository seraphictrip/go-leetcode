package trie

type Trie struct {
	children map[rune]*Trie
	isWord   bool
}

func Constructor() Trie {
	return Trie{
		children: make(map[rune]*Trie),
	}
}

func NewTrie() *Trie {
	node := Constructor()
	return &node
}

func (t *Trie) Insert(word string) {
	cur := t
	for _, ch := range word {
		if node, ok := cur.children[ch]; ok {
			cur = node
		} else {
			node = NewTrie()
			cur.children[ch] = node
			cur = node
		}
	}
	cur.isWord = true
}

func (t *Trie) Search(word string) bool {
	cur := t
	for _, ch := range word {
		node, ok := cur.children[ch]
		if !ok {
			return false
		}
		cur = node
	}
	return cur.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	cur := t
	for _, ch := range prefix {
		node, ok := cur.children[ch]
		if !ok {
			return false
		}
		cur = node
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
