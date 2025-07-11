package main

import (
	"fmt"
)

func pipeline() {
	/*
	   A pipeline is a sequence of stages connected by channels. The output of one stage is the input of the next. This pattern is great for streaming data processing where you want to perform a series of transformations.

	   The Pattern:

	   A "source" stage creates a channel and sends initial data to it.

	   A "processor" stage reads from the source channel, processes the data, and sends the results to a new channel.

	   A "sink" or "consumer" stage reads from the final channel.
	*/
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	done := make(chan struct{})

	in := gen(done, nums...)
	out := square(done, in)

	for n := range out {
		fmt.Println(n)
	}

}

func gen(done chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func square(done chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}
