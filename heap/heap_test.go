package heap

import (
	"testing"

	"github.com/houshengbo/golang-skills/util"
)

func TestHeap(t *testing.T) {
	heap := &MaxHeap{}

	values := []int{2, 1, 3}
	for _, v := range values {
		heap.Insert(v)
	}

	util.AssertEqual(t, len(heap.array), 3)
	util.AssertEqual(t, heap.array[0], 3)
}
