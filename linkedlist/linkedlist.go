package linkedlist

import "fmt"

// A linked list consists of nodes where each node contains a data field and a reference(link)
// to the next node in the list. O(n).

// It is fast to add or remove values at the beginning of the list. O(1)
// For an array, adding a value in front can be very time-costing.

// A node consists of the data and the pointers pointing to the next node.
type node struct {
	data int
	next *node
}

// A linked list only holds the head, which points to the first node.
type linkedlist struct {
	head   *node
	length int
}

func (l *linkedlist) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedlist) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedlist) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}
