package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		// Do some work here...
		results <- j * 2
		fmt.Println("Worker", id, "finished job", j)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	// Create channels for passing jobs and results
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Start up some workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
		time.Sleep(2 * time.Second)
	}

	// Send some jobs to the workers
	for j := 1; j <= 5; j++ {
		jobs <- j
		time.Sleep(2 * time.Second)
	}
	close(jobs)

	// Collect the results from the workers
	for a := 1; a <= 5; a++ {
		<-results
		time.Sleep(2 * time.Second)
	}
}
