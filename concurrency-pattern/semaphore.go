package concurrencypattern

import (
	"fmt"
	"sync"
	"time"
)

// Simulate downloading a file
func downloadFile(fileID int, semaphore chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	// Acquire the semaphore (take up a slot)
	semaphore <- struct{}{}

	// Simulate the download process
	fmt.Printf("Downloading file %d...\n", fileID)
	time.Sleep(2 * time.Second) // Simulating time taken to download
	fmt.Printf("Finished downloading file %d.\n", fileID)

	// Release the semaphore (free up a slot)
	<-semaphore
}

func SemaphorePattern() {
	// Create a semaphore channel that allows only 3 goroutines at a time
	semaphore := make(chan struct{}, 3)
	var wg sync.WaitGroup
	// Simulate downloading 10 files
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go downloadFile(i, semaphore, &wg)
	}

	wg.Wait()
	fmt.Println("Finished downloading!")
}
