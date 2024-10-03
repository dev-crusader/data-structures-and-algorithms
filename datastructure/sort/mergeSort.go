package sort

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

func SortWithMerge() {
	// inputArr := []int{50, -8, -96, -63, -73, 95, -69, 16, -38, 53, 72, -71, -92, 25, 59}
	inputArr := generateRandomArray(-100, 100, 200)
	arr := make([]int, len(inputArr))
	copy(arr, inputArr)
	// fmt.Printf("Before:\t%v\n", inputArr)
	t1 := time.Now()
	mergeSort(inputArr, 0, len(inputArr)-1)
	// fmt.Printf("After:\t%v\n", inputArr)
	fmt.Printf("Took %d nanoseconds\n", time.Since(t1).Nanoseconds())

	// Merge sort implementation concurrently splitting and sorting
	t2 := time.Now()

	var wg sync.WaitGroup
	wg.Add(1)
	go mergeSortConcurrent(arr, 0, len(arr)-1, &wg)
	wg.Wait()
	// fmt.Printf("After:\t%v\n", arr)
	fmt.Printf("Took %d nanoseconds\n", time.Since(t2).Nanoseconds())
	fmt.Println("Done")
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

func generateRandomArray(min, max, total int) []int {
	rand.Seed(uint64(time.Now().UnixNano()))
	var num []int

	for i := 0; i < total; i++ {
		rn := rand.Intn(max-min) + min
		num = append(num, rn)
	}
	return num
}

func mergeSortConcurrent(arr []int, left, right int, wg *sync.WaitGroup) {
	defer wg.Done()
	if right <= left {
		return
	}

	mid := (left + right) / 2

	var leftWg, rightWg sync.WaitGroup

	leftWg.Add(1)
	rightWg.Add(1)
	go mergeSortConcurrent(arr, left, mid, &leftWg)
	go mergeSortConcurrent(arr, mid+1, right, &rightWg)
	leftWg.Wait()
	rightWg.Wait()
	mergeAllStep(arr, left, mid, right)
}
