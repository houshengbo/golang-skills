package binarytree

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(key int) {
	if t.Root == nil {
		t.Root = &Node{Key: key}
	} else {
		t.Root.Insert(key)
	}
}

func (t *Tree) Search(key int) bool {
	if t.Root == nil {
		return false
	} else {
		return t.Root.Search(key)
	}
}

// Each node has a key and poitns to a left node and a right node. Left node is smaller, and the right node
// is larger. Node without children is called leave.
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (node *Node) Insert(k int) {
	if node.Key < k {
		if node.Right == nil {
			node.Right = &Node{Key: k}
		} else {
			node.Right.Insert(k)
		}
	} else if node.Key >= k {
		if node.Left == nil {
			node.Left = &Node{Key: k}
		} else {
			node.Left.Insert(k)
		}
	}
}

func (node *Node) Search(k int) bool {
	if node == nil {
		return false
	}
	if node.Key < k {
		return node.Right.Search(k)
	} else if node.Key > k {
		return node.Left.Search(k)
	}
	return true
}
