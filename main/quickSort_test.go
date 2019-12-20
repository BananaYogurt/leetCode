package main

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	testArray := []int{7, 9, 2, 23, 1, 0, 7, 12, 33, 28, 55, 128, -2, 66, -4, 37}
	QuickSort(testArray, 0, len(testArray)-1)
	for i := range testArray {
		fmt.Print(testArray[i], ",")
	}
}
