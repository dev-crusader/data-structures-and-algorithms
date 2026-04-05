package sort

// BubbleSort sorts an integer slice using the bubble sort algorithm
func BubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        // Last i elements are already in place
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                // Swap arr[j] and arr[j+1]
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}