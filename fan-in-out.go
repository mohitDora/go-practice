package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func producer2(jobs chan<- int, numbers []int) {
	for _, number := range numbers {
		jobs <- number
	}
	close(jobs)
}

func worker2(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		results <- job * 2
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}
}

func fan_in(resultsChannel ...chan int) <-chan int {
	var wg sync.WaitGroup
	finalResults := make(chan int)

	for _, ch := range resultsChannel {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for result := range c {
				finalResults <- result
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(finalResults)
	}()

	return finalResults
}

func fan_in_out() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	jobs := make(chan int, len(numbers))
	results := make([]chan int, 0, len(numbers))
	go producer2(jobs, numbers)
	var wg sync.WaitGroup
	for i := 0; i < len(numbers); i++ {
		results = append(results, make(chan int))
		wg.Add(1)
		go worker2(i, jobs, results[i], &wg)
	}

	go func() {
		wg.Wait()
		for _, ch := range results {
			close(ch)
		}
	}()
	finalResults := fan_in(results...)
	fmt.Println("\nMain: Collecting results...")
	for result := range finalResults {
		fmt.Println("Collected result:", result)
	}

	fmt.Println("\nMain: All tasks completed.")
}
