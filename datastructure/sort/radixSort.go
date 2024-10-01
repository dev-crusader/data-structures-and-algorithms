package sorting

import (
	"fmt"
	"math"
	"sort"
)

func Radix() {
	inputArr := []int{6151, 4256, 4485, 3634, 4378, 5075, 3407, 7189, 4211, 7561, 5854, 2610, 6476, 3491, 7645}

	fmt.Println("\nBefore: \n", inputArr)
	radixSort(inputArr, 4, 10)
	fmt.Println("\nAfter: \n", inputArr)
}

func radixSort(arr []int, width, radix int) {
	for i := 0; i < width; i++ {
		// sortStable(arr, float64(i), radix)
		countingSort(arr, i, radix)
	}
}

func countingSort(arr []int, p, radix int) {
	countArr := make([]int, radix)
	for k := range arr {
		countArr[getIndex(arr[k], p, radix)]++
	}
	sum := 0
	for i, k := range countArr {
		sum += k
		countArr[i] = sum
	}
	shuffleArray(arr, countArr, p, radix)
}

func shuffleArray(arr, ctArr []int, p, r int) {
	temp := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		index := getIndex(arr[i], p, r)
		ctArr[index] = ctArr[index] - 1
		temp[ctArr[index]] = arr[i]
	}
	copy(arr, temp)
}

func getIndex(num, pos, radix int) int {
	return (num / int(math.Pow(10.0, float64(pos))) % radix)
}

func sortStable(arr []int, position float64, radix int) {
	sort.SliceStable(arr, func(i, j int) bool {
		return ((arr[i] / int(math.Pow(10.0, position))) % radix) < ((arr[j] / int(math.Pow(10.0, position))) % radix)
	})
}
