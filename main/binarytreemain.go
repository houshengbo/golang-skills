package main

import (
	"fmt"

	"github.com/houshengbo/golang-skills/binarytree"
)

func main() {
	tree := &binarytree.Node{Key: 100}
	tree.Insert(200)
	tree.Insert(300)

	fmt.Println(tree)
	fmt.Println(tree.Search(1))
	fmt.Println(tree.Search(200))
}
