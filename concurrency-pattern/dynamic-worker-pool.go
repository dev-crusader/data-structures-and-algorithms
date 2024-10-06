package concurrencypattern

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	id int
}

// Worker pool struct that can resize dynamically
type WorkerPool struct {
	taskQueue   chan Task      // Task channel
	workerCount int            // Number of active workers
	maxWorkers  int            // Maximum workers allowed
	mu          sync.Mutex     // Mutex to handle resizing
	wg          sync.WaitGroup // WaitGroup to track completion
}

// NewWorkerPool initializes a new worker pool with a task queue
func NewWorkerPool(initialWorkers, maxWorkers int) *WorkerPool {
	return &WorkerPool{
		taskQueue:   make(chan Task, 10), // Buffer for incoming tasks
		workerCount: initialWorkers,
		maxWorkers:  maxWorkers,
	}
}

// Start initializes the worker pool and starts the initial workers
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workerCount; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

// Submit adds a new task to the task queue
func (wp *WorkerPool) Submit(task Task) {
	wp.taskQueue <- task
	wp.dynamicResize()
}

// Worker simulates task processing by workers
func (wp *WorkerPool) worker(workerID int) {
	defer wp.wg.Done()
	fmt.Printf("Worker %d started.\n", workerID)

	for task := range wp.taskQueue {
		fmt.Printf("Worker %d processing task %d\n", workerID, task.id)
		time.Sleep(2 * time.Second) // Simulate work
		fmt.Printf("Worker %d finished task %d\n", workerID, task.id)
	}

	fmt.Printf("Worker %d shutting down.\n", workerID)
}

// dynamicResize adjusts the number of workers based on task volume
func (wp *WorkerPool) dynamicResize() {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	// If there are more tasks than workers, increase worker count
	if len(wp.taskQueue) > wp.workerCount && wp.workerCount < wp.maxWorkers {
		fmt.Println("Increasing worker count...")
		wp.workerCount++
		wp.wg.Add(1)
		go wp.worker(wp.workerCount)
	}
}

// Shutdown gracefully closes the worker pool
func (wp *WorkerPool) Shutdown() {
	close(wp.taskQueue) // Close task channel, stopping workers
	wp.wg.Wait()        // Wait for all workers to finish processing
	fmt.Println("All workers shut down.")
}

func DynamicWorkerPool() {
	// Initialize a worker pool with 2 initial workers and a maximum of 5 workers
	workerPool := NewWorkerPool(2, 5)
	workerPool.Start()

	// Simulate submitting 10 tasks to the worker pool
	for i := 1; i <= 10; i++ {
		task := Task{id: i}
		workerPool.Submit(task)
		time.Sleep(500 * time.Millisecond) // Simulate task submission over time
	}

	// Give some time for processing before shutting down
	time.Sleep(5 * time.Second)
	workerPool.Shutdown()
}
