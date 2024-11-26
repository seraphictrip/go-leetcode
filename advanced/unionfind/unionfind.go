package unionfind

type UnionFind[T comparable] struct {
	parent map[T]T
	rank   map[T]int
}
