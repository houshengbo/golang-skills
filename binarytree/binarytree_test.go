package binarytree

import (
	"testing"

	"github.com/houshengbo/golang-skills/util"
)

func TestBinaryTree(t *testing.T) {
	tree := &Tree{}
	tree.Insert(76)
	tree.Insert(30)
	tree.Insert(78)
	tree.Insert(63)
	tree.Insert(98)
	tree.Insert(760)
	tree.Insert(60)
	tree.Insert(39)
	tree.Insert(23)
	tree.Insert(375)

	util.AssertEqual(t, tree.Search(23), true)
	util.AssertEqual(t, tree.Search(375), true)
	util.AssertEqual(t, tree.Search(64), false)
	util.AssertEqual(t, tree.Search(60), true)

	util.AssertEqual(t, tree.Search(378), false)
	util.AssertEqual(t, tree.Search(239), false)
}
