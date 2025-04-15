package linkedlist

type Node[T comparable] struct {
	val  T
	next *Node[T]
	prev *Node[T]
}

func NewNode[T comparable](val T) *Node[T] {
	return &Node[T]{val, nil, nil}
}
