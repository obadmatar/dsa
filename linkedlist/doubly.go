package linkedlist

type Doubly[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewDoubly[T comparable]() *Doubly[T] {
	return new(Doubly[T])
}

// Append adds val at the end of the list
func (ll *Doubly[T]) Append(val T) {
	n := NewNode(val)
	if ll.Empty() {
		ll.head = n
		ll.tail = n
	} else {
		n.prev = ll.tail
		ll.tail.next = n
		ll.tail = n
	}
	ll.size++
}

// Prepend adds val at the start of the list
func (ll *Doubly[T]) Prepend(val T) {
	n := NewNode(val)
	if ll.Empty() {
		ll.head = n
		ll.tail = n
	} else {
		n.next = ll.head
		ll.head.prev = n
		ll.head = n
	}
	ll.size++
}

// Get returns node value at the given index
func (ll *Doubly[T]) Get(index int) (T, bool) {
	if !ll.withinRange(index) {
		var t T
		return t, false
	}

	if index < ll.size/2 {
		n := ll.head
		for i := 0; i != index; i++ {
			n = n.next
		}
		return n.val, true
	}

	n := ll.tail
	for i := ll.size - 1; i != index; i-- {
		n = n.prev
	}
	return n.val, true
}

// Remove removes node value at the given index
func (ll *Doubly[T]) Remove(index int) {
	if !ll.withinRange(index) {
		return
	}

	node := ll.head
	for i := 0; i != index; i++ {
		node = node.next
	}

	if node.prev != nil {
		node.prev.next = node.next
	} else {
		ll.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		ll.tail = node.prev
	}

	ll.size--
}

// Contains checks if val is present in the list
func (ll *Doubly[T]) Contains(val T) bool {
	for n := ll.head; n != nil; n = n.next {
		if n.val == val {
			return true
		}
	}
	return false
}

// Values returns all values in the list
func (ll *Doubly[T]) Values() []T {
	values := make([]T, ll.size)
	for i, n := 0, ll.head; n != nil; i, n = i+1, n.next {
		values[i] = n.val
	}
	return values
}

// Empty returns true if list does not contain any nodes
func (ll *Doubly[T]) Empty() bool {
	return ll.size == 0
}

// Size returns number of nodes within the list
func (ll *Doubly[T]) Size() int {
	return ll.size
}

// withinRange checks that index is within list size
func (ll *Doubly[T]) withinRange(index int) bool {
	return index >= 0 && index < ll.size
}
