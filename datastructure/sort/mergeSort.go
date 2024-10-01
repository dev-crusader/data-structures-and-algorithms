package sorting

import "fmt"

func SortWithMerge() {
	inputArr := []int{50, -8, -96, -63, -73, 95, -69, 16, -38, 53, 72, -71, -92, 25, 59}
	fmt.Printf("Before:\t%v\n", inputArr)
	mergeSort(inputArr, 0, len(inputArr)-1)
	fmt.Printf("After:\t%v\n", inputArr)
}
func mergeSort(arr []int, start, end int) {
	if end <= start {
		return
	}

	mid := (start + end) / 2

	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	mergeAllStep(arr, start, mid, end)
}

func mergeAllStep(inpArr []int, start, mid, end int) {
	if inpArr[mid] <= inpArr[mid+1] {
		return
	}
	tempArr := make([]int, end-start+1)
	i, j, tempIndex := start, mid+1, 0
	for i <= mid && j <= end {
		if inpArr[i] < inpArr[j] {
			tempArr[tempIndex] = inpArr[i]
			i++
		} else {
			tempArr[tempIndex] = inpArr[j]
			j++
		}
		tempIndex++
	}
	for i <= mid {
		tempArr[tempIndex] = inpArr[i]
		i++
		tempIndex++
	}
	for j <= end {
		tempArr[tempIndex] = inpArr[j]
		j++
		tempIndex++
	}
	k := start
	for _, v := range tempArr {
		inpArr[k] = v
		k++
	}
}
