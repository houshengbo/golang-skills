package hashmap

import (
	"fmt"
	"testing"

	"github.com/houshengbo/golang-skills/util"
)

func TestHashMap(t *testing.T) {
	testHashTable := Init()
	fmt.Println(testHashTable)
	fmt.Println(hash("randy"))

	testHashTable.Insert("randy")
	testHashTable.Insert("eric")
	testHashTable.Insert("kyle")
	testHashTable.Insert("tim")
	testHashTable.Insert("token")

	util.AssertEqual(t, testHashTable.Search("eric"), true)
	util.AssertEqual(t, testHashTable.Search("randy"), true)
	util.AssertEqual(t, testHashTable.Search("test"), false)
}
