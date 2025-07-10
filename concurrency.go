package main

import (
	"fmt"
	"sync"
	"time"
)

func concurrency() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		greet("John")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		greet("Harry")
	}()

	/* channels - used to communicate between goroutines

		a channel can only send or receive a value of a specific type eg int, string, struct, etc
		By default sending and receiving operations are blocking in nature
		a send operation will block until there is a corresponding receive operation
		a receive operation will block until there is a corresponding send operation

	Unbuffered channel -  has a capacity of 0 and only allows communication when both sender and receiver are ready

	Buffered channel - has a specific capacity and allows sender to send values to the channel until it is full and receiver to receive values from the channel until it is empty w/o blocking
	*/

	// unbuffered channel
	channel := make(chan string)

	wg.Add(1)
	go func() {
		defer wg.Done()
		worker(1, channel)
	}()

	fmt.Println("Sending message to worker")
	time.Sleep(5 * time.Second)
	channel <- "Hello"
	fmt.Println("Message sent to worker")

	// buffered channel
	channel2 := make(chan string, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		sender(channel2)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		receiver(channel2)
	}()

	// producer consumer pattern
	channel3 := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		producer(channel3)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		consumer(channel3)
	}()

	wg.Wait()

	// select statement - used to wait on multiple channel operations
	channel4 := make(chan string)
	channel5 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel4 <- "Hello"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		channel5 <- "World"
	}()

	select {
	case msg1 := <-channel4:
		fmt.Println(msg1)
	case msg2 := <-channel5:
		fmt.Println(msg2)
	default:
		fmt.Println("No message ready")
	}

}

func greet(name string) {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello", name)
}

func worker(id int, channel chan string) {
	fmt.Println("Worker", id, "waiting for message")
	msg := <-channel
	fmt.Println("Worker", id, "received message", msg)
}

func sender(channel chan string) {
	channel <- "1"
	fmt.Println("Message sent to receiver")
	channel <- "2"
	fmt.Println("Message sent to receiver")
	channel <- "3"
	fmt.Println("Message sent to receiver")
}

func receiver(channel chan string) {
	time.Sleep(3 * time.Second)
	msg := <-channel
	fmt.Println("Message received from sender", msg)
	msg = <-channel
	fmt.Println("Message received from sender", msg)
	msg = <-channel
	fmt.Println("Message received from sender", msg)
}

func producer(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel)
}

func consumer(channel chan int) {
	for i := range channel {
		fmt.Println(i)
	}
}
