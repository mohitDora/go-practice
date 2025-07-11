package main

import (
	"fmt"
	"sync"
	"time"
)

func workerFunction(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2
	}
}

func workerPool() {

	/*
	   A worker pool is a design pattern used to limit the number of goroutines that handle a set of tasks. This is essential for preventing resource exhaustion (e.g., too many database connections, too many open files) when you have a large number of jobs to process.

	   The Pattern:

	   Create a channel for jobs.

	   Create a fixed number of "worker" goroutines.

	   Each worker listens for jobs on the channel.

	   The main goroutine sends jobs to the channel.

	   Once all jobs are sent, the jobs channel is closed to signal to the workers that no more jobs are coming.
	*/

	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			workerFunction(id, jobs, results)
		}(w)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
	close(results)

	for r := range results {
		fmt.Println("Result", r)
	}
}
