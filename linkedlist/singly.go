package linkedlist

type Singly[T comparable] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

func NewSingly[T comparable]() *Singly[T] {
	// TODO: accept list of values to be added to the list at initalization
	return new(Singly[T])
}

// Append adds val at the end of the list
func (ll *Singly[T]) Append(val T) {
	n := NewNode(val)
	if ll.size == 0 {
		ll.head = n
		ll.tail = n
	} else {
		ll.tail.next = n
		ll.tail = n
	}
	ll.size++
}

// Prepend adds val at the start of the list
func (ll *Singly[T]) Prepend(val T) {
	n := NewNode(val)
	n.next = ll.head

	ll.head = n

	if ll.size == 0 {
		ll.tail = n
	}

	ll.size++
}

// Get returns node value at the given index
func (ll *Singly[T]) Get(index int) (T, bool) {
	if !ll.withinRange(index) {
		var t T
		return t, false
	}

	node := ll.head
	for i := 0; i != index; i++ {
		node = node.next
	}

	return node.val, true
}

// Remove removes node value at the given index
func (ll *Singly[T]) Remove(index int) {
	if !ll.withinRange(index) {
		return
	}

	var prev, curr *Node[T]
	prev, curr = nil, ll.head

	for i := 0; i != index; i++ {
		prev = curr
		curr = curr.next
	}

	if curr == ll.head {
		ll.head = curr.next
	}

	if curr == ll.tail {
		ll.tail = prev
	}

	if prev != nil {
		prev.next = curr.next
	}

	curr = nil

	ll.size--
}

// Contains checks if val is present in the list
func (ll *Singly[T]) Contains(val T) bool {
	for n := ll.head; n != nil; n = n.next {
		if n.val == val {
			return true
		}
	}
	return false
}

// Values returns all values in the list
func (ll *Singly[T]) Values() []T {
	values := make([]T, ll.size)
	for i, n := 0, ll.head; n != nil; i, n = i+1, n.next {
		values[i] = n.val
	}
	return values
}

// Empty returns true if list does not contain any nodes
func (ll *Singly[T]) Empty() bool {
	return ll.size == 0
}

// Size returns number of nodes within the list
func (ll *Singly[T]) Size() int {
	return ll.size
}

// withinRange checks that index is within list size
func (ll *Singly[T]) withinRange(index int) bool {
	return index >= 0 && index < ll.size
}
