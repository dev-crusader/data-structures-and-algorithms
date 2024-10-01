package sorting

import "fmt"

func QuickSortApproach() {
	inputArr := []int{50, -8, -96, -63, -73, 95, -69, 16, -38, 53, 72, -71, -92, 25, 59}
	fmt.Printf("Before:\t%v\n", inputArr)
	quickSort(inputArr, 0, len(inputArr)-1)
	fmt.Printf("After:\t%v\n", inputArr)
}

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	pivotIndex := partition(arr, start, end)
	quickSort(arr, start, pivotIndex-1)
	quickSort(arr, pivotIndex+1, end)
}

func partition(arr []int, start, end int) int {
	pivotElement := arr[end]
	i := start - 1

	for j := start; j < end; j++ {
		if arr[j] < pivotElement {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[end] = arr[i+1]
	arr[i+1] = pivotElement

	return i + 1
}
