package main

import (
	"fmt"
)

func quickSortPivotFirstElement(array []int) []int {
	quickSortRecPivotFirstElement(array, 0, len(array)-1)
	return array
}

func quickSortRecPivotFirstElement(array []int, startIdx int, endIdx int) {
	// base case: contain a singe element or empty subarray
	if startIdx >= endIdx {
		return
	}

	// 1. pick the first element as a pivot
	pivotIdx := startIdx
	leftIdx := startIdx + 1
	rightIdx := endIdx

	// 2. rearrange the array
	// 2.1 move elements smaller than the pivot to the left. Otherwise, move to the right
	for leftIdx <= rightIdx {
		if array[leftIdx] > array[pivotIdx] && array[pivotIdx] > array[rightIdx] {
			swap(leftIdx, rightIdx, array)
		}
		// equal sign for duplicate numbers
		if array[leftIdx] <= array[pivotIdx] {
			leftIdx += 1
		}
		if array[rightIdx] >= array[pivotIdx] {
			rightIdx -= 1
		}
	}
	// 2.2 place the pivot at the current position in the sorted array
	swap(pivotIdx, rightIdx, array)

	// current array status: left subarray | original pivot value | right subarray
	// 3. partition
	quickSortRecPivotFirstElement(array, startIdx, rightIdx-1)
	quickSortRecPivotFirstElement(array, rightIdx+1, endIdx)
}

func swap(i, j int, array []int) {
	array[i], array[j] = array[j], array[i]
}

func main() {
	array := []int{8, 5, 2, 4, 11, 9, 22}
	fmt.Println("pick the 1st element as a pivot", quickSortPivotFirstElement(array))
}

/* output
pick the 1st element as a pivot [2 4 5 8 9 11 22]
*/
