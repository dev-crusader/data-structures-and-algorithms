package concurrencypattern

import "fmt"

func stage1(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2 // Process data
		}
		close(out)
	}()
	return out
}

func stage2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n + 1 // Further process data
		}
		close(out)
	}()
	return out
}

func Pipeline() {
	in := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			in <- i // Send values
		}
		close(in) // Close input channel
	}()

	// Create a pipeline
	out := stage2(stage1(in))

	for result := range out {
		fmt.Println(result) // Output results
	}
}
