package concurrencypattern

import (
	"fmt"
	"sync"
)

// Worker function that processes a task
func workerRoutine(task int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Processed task: %d\n", task)
	results <- task * 2 // Send the processed result to the results channel
}

func FanInFanOut() {
	var wg sync.WaitGroup
	tasks := []int{1, 2, 3, 4, 5}         // Example tasks (each task is just an int)
	results := make(chan int, len(tasks)) // Buffered channel for results

	// Fan-out: Create a new goroutine for each task
	for _, task := range tasks {
		wg.Add(1)
		go workerRoutine(task, results, &wg)
	}

	// Fan-in: Wait for all workers to finish and then close the results channel
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect and print the results
	for result := range results {
		fmt.Printf("Collected result: %d\n", result)
	}
}
