package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	data map[string]string
}

func (c *Cache) Read(key string, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	c.RLock()
	defer c.RUnlock()
	time.Sleep(2 * time.Second)
	value, ok := c.data[key]
	if !ok {
		fmt.Printf("Reader %d: Key %s not found\n", id, key)
		return
	}
	fmt.Printf("Reader %d: Key %s, Value %s\n", id, key, value)
}

func (c *Cache) Write(key, value string, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	c.Lock()
	defer c.Unlock()
	fmt.Printf("Writer %d: Writing Key %s, Value %s\n", id, key, value)
	time.Sleep(1 * time.Second)
	c.data[key] = value
}

var cache = Cache{data: make(map[string]string)}

func rwMutex() {
	/*
	   - RWMutex is a reader/writer mutual exclusion lock. The lock can be held by an arbitrary number of readers or a single writer. But only one writer can hold the lock at a time.
	*/
	wg := sync.WaitGroup{}
	wg.Add(2)
	go cache.Write("user1", "Alice", 1, &wg)
	go cache.Write("user2", "Bob", 2, &wg)
	wg.Wait()

	fmt.Println("Initial writes complete. Starting concurrent reads and a new write.")

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go cache.Read("user1", i, &wg)
	}
	wg.Add(1)
	go cache.Write("user3", "Charlie", 3, &wg)

	wg.Wait()
	fmt.Println("All operations finished.")

}
