package linkedlist

import (
	"testing"
)

func TestLinkedlist(t *testing.T) {
	myList := linkedlist{}
	node1 := &node{data: 48}
	node2 := &node{data: 18}
	node3 := &node{data: 16}
	node4 := &node{data: 11}
	node5 := &node{data: 7}
	node6 := &node{data: 2}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)

	myList.printListData()
	myList.deleteWithValue(16)
	myList.printListData()

	myList.deleteWithValue(186)
	myList.printListData()

	emptyList := linkedlist{}
	emptyList.deleteWithValue(12)
}
