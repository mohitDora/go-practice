package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

func pool_() {
	/*
		sync.Pool is designed for managing a temporary pool of objects that are expensive to create. It is a highly specialized tool used to reduce pressure on the garbage collector by reusing objects that would otherwise be discarded.
	*/

	wg := sync.WaitGroup{}

	fmt.Println("--- PHASE 1: Initial creation ---")
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go processRequests(i, &wg)
	}

	wg.Wait()
	fmt.Println("\n--- PHASE 2: Reusing existing buffers ---")
	for i := 6; i <= 10; i++ {
		wg.Add(1)
		go processRequests(i, &wg)
	}

	wg.Wait()
	fmt.Println("\nAll requests processed.")
}

const bufferSize = 1024 * 1024

var bufferPool = sync.Pool{
	New: func() interface{} {
		fmt.Println("Creating new buffer")
		return new(bytes.Buffer)
	}, // New is called when the pool is empty and a new object is needed
}

func processRequests(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	buffer := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buffer)
	time.Sleep(1 * time.Second)
	fmt.Printf("Request %d processed\n", id)
}
