package main

import (
	"fmt"
	"sync"
)

var count int
var mu sync.Mutex

func mutex_() {
	count = 0
	wg := sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
	fmt.Printf("Counter: %v\n", count)
}

func increment() {
	for i := 0; i < 1000; i++ {
		mu.Lock()
		count++
		mu.Unlock()
	}
}
