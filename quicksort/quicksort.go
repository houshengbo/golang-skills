package quicksort

import (
	"math/rand"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	arr, left := partition(arr)
	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}

func partition(arr []int) ([]int, int) {
	// Any value in the array can become the pivot
	pivot := getPivotIndex(arr)

	left, right := 0, len(arr)-1
	arr[pivot], arr[right] = arr[right], arr[pivot]

	for j := left; j < right; j++ {
		if arr[j] < arr[right] {
			// Swap the value at left(larger than pivot) with the value at j(smaller than pivot)
			// It is possible that left and j are the same
			arr[left], arr[j] = arr[j], arr[left]
			// left is used to mark the possible value that is larger than pivot
			// left is already smaller than or equal to j
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]
	return arr, left
}

func getPivotIndex(arr []int) int {
	return rand.Int() % len(arr)
}
