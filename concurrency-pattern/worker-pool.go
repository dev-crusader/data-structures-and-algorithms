package concurrencypattern

import (
	"fmt"
	"sync"
)

// Worker function that processes input numbers
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// Simulate some work
		result := job * job // Example processing (squaring the number)
		fmt.Printf("Worker %d processed job %d; result = %d\n", id, job, result)
		results <- result
	}
}

func RunWorkerPool() {
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// Start worker goroutines (fan-out)
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs to the jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Close jobs channel as no more jobs will be sent

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results) // Close results channel after workers finish
	}()
	// Collect results
	for result := range results {
		fmt.Printf("Collected result: %d\n", result)
	}
}
