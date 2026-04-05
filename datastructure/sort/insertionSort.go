package sort

// InsertionSort sorts an integer slice using the insertion sort algorithm
func InsertionSort(arr []int) {
    n := len(arr)
    for i := 1; i < n; i++ {
        key := arr[i]
        j := i - 1

        // Move elements of arr[0..i-1], that are greater than key, to one position ahead
        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j--
        }
        // Place the key at after the element just smaller than it.
        arr[j+1] = key
    }
}