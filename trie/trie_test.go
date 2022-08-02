package trie

import (
	"github.com/houshengbo/golang-skills/util"
	"testing"
)

func TestTrie(t *testing.T) {
	myTrie := InitTrie()
	toAdd := []string{
		"hajkfhs",
		"ffds",
		"fsfs",
		"ewefegefw",
		"fsfgge",
		"qwdqdwf",
	}

	for _, v := range toAdd {
		myTrie.Insert(v)
	}

	util.AssertDeepEqual(t, myTrie.Search("ewefwegefw"), false)
	util.AssertDeepEqual(t, myTrie.Search("ewefegefw"), true)
}
