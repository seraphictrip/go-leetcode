package arr

import "slices"

/*
Array Implementation of a binary tree
self = i
leftchild = 2i+1
rightchild = 2i+2
*/

// [R, C1, C2, C1_1, C1_2, C2_1, C2_2]

type BinaryTree[T any] struct {
	internal []T
}

func (bt *BinaryTree[T]) Insert(item T) {
	// for a balanced tree, the next insertion point should just be next index right?
	// explore
	bt.internal = append(bt.internal, item)
}

func (bt *BinaryTree[T]) ToArray() []T {
	return slices.Clone(bt.internal)
}
