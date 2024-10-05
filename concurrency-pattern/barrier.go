package concurrencypattern

import (
	"fmt"
	"sync"
	"time"
)

// Function to simulate a multi-stage data processing worker
func processData(id int, wg *sync.WaitGroup, barrier chan struct{}) {
	defer wg.Done()

	// Stage 1: Fetch data
	fmt.Printf("Worker %d is fetching data...\n", id)
	time.Sleep(time.Second * 2) // Simulating data fetching
	fmt.Printf("Worker %d finished fetching data.\n", id)

	// Wait at barrier for other workers to finish fetching data
	barrier <- struct{}{} // Signal that this worker has reached the barrier
	fmt.Printf("Worker %d waiting at the barrier after fetching.\n", id)

	// Stage 2: Transform data (After all workers have fetched data)
	fmt.Printf("Worker %d is transforming data...\n", id)
	time.Sleep(time.Second * 1) // Simulating data transformation
	fmt.Printf("Worker %d finished transforming data.\n", id)

	// Wait at barrier for other workers to finish transforming data
	barrier <- struct{}{} // Signal that this worker has reached the barrier
	fmt.Printf("Worker %d waiting at the barrier after transforming.\n", id)

	// Stage 3: Write data (After all workers have transformed data)
	fmt.Printf("Worker %d is writing data to the database...\n", id)
	time.Sleep(time.Second * 1) // Simulating writing data
	fmt.Printf("Worker %d finished writing data.\n", id)
}

func Barrier() {
	numWorkers := 3
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Create a barrier (channel) to signal when all workers reach it
	barrier := make(chan struct{}, numWorkers*2) // Two barriers (one after fetching, one after transforming)

	// Start multiple workers (goroutines)
	for i := 1; i <= numWorkers; i++ {
		go processData(i, &wg, barrier)
	}

	// Wait for all workers to reach the first barrier (after fetching data)
	for i := 0; i < numWorkers; i++ {
		<-barrier // Wait for each worker to signal they have finished fetching
	}
	fmt.Println("All workers finished fetching data. Moving to the next stage.")

	// Wait for all workers to reach the second barrier (after transforming data)
	for i := 0; i < numWorkers; i++ {
		<-barrier // Wait for each worker to signal they have finished transforming
	}
	fmt.Println("All workers finished transforming data. Proceeding to the final stage.")

	// Wait for all workers to finish processing
	wg.Wait()
	fmt.Println("All workers finished writing data. Processing completed.")
}
