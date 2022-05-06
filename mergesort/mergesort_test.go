package mergesort

import (
	"testing"

	"github.com/houshengbo/golang-skills/util"
)

func TestMergeSort(t *testing.T) {
	slice := []int{100, 2, 6, 3, 8, 5, 9, 4, 0}
	expected := []int{0, 2, 3, 4, 5, 6, 8, 9, 100}
	result := mergeSort(slice)
	util.AssertDeepEqual(t, result, expected)
}
