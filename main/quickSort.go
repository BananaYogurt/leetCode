package main

//快排

func QuickSort(array []int, left int, right int) {
	start := left
	end := right

	if left < right {
		tmpV := array[start]
		for start < end {
			for start < end && array[end] >= tmpV {
				end--
			}
			array[start] = array[end]
			for end > start && array[start] <= tmpV {
				start++
			}
			array[end] = array[start]
		}
		array[start] = tmpV
		QuickSort(array, left, start-1)
		QuickSort(array, start+1, right)
	}
}
