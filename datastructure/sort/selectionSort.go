package sort

func SelectionSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        // Assume the current position contains the minimum
        minIdx := i
        // Find the minimum element in the unsorted part
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIdx] {
                minIdx = j
            }
        }
        // Swap the found minimum element with the first element of the unsorted part
        arr[i], arr[minIdx] = arr[minIdx], arr[i]
    }
}