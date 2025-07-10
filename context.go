package main

import (
	"context"
	"fmt"
	"time"
)

func context_() {
	/*
					   Context is a package that provides a way to pass request-scoped values, deadlines, and cancellation signals across API boundaries to all the goroutines involved in handling a request.

					   suppose a web reuest invoke multiple goroutines, if the request is cancelled, we want to
					   - cancel all the goroutines that are still running
					   -set a deadline for the request
					   - pass request-scoped values to all the goroutines

					   All these can be done using context package

					   Context is an interface that has four methods:
					   - Deadline() (deadline time.Time, ok bool)
					   - Done() <-chan struct{} returns a channel that is closed when the context is cancelled

					   - Err() error returns the reason for cancellation
					   - Value(key interface{}) interface{} returns the value associated with this context for key, or nil if no value is associated with key

					   Creating context :
					   - context.Background() : returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline. It is typically used in main, init, and tests, and as the top-level Context for incoming requests.
					   - context.TODO() : returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline. It is typically used when a function requires a Context but the caller does not have one.

					   There are several functions in the context package :
					   - WithCancel(parent Context) (ctx Context, cancel CancelFunc) : returns a copy of parent with a new Done channel. The returned context's Done channel is closed when the returned cancel function is called or when the parent context's Done channel is
					   closed, whichever happens first.

				      - WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) : returns a new context and a cancel function. The context is automatically canceled after the timeout duration elapses or when the cancel is explicitly called.

				      - WithDeadLine(parent Context, deadline time.Time) (Context, CancelFunc) :
				      similar to WithTimeout, but the deadline is specified as a time.Time instead of a duration.

		             - WithValue(parent Context, key, val interface{}) Context : returns a copy of parent with a new value associated with key.
	*/

	// with cancel
	ctxCancel, cancel := context.WithCancel(context.Background())

	simulateRequest(ctxCancel)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("====Timeout Example====")

	// with timeout
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelTimeout()
	simulateRequest(ctxTimeout)
	time.Sleep(3 * time.Second)
}

func performTask(ctx context.Context, task string, delay int) {
	select {
	case <-time.After(time.Duration(delay) * time.Second):
		fmt.Println("Task", task, "Done")
	case <-ctx.Done():
		fmt.Printf("%s: Canceled! Reason: %v\n", task, ctx.Err())
	}
}

func simulateRequest(ctx context.Context) {
	go performTask(ctx, "DB Query", 1)
	go performTask(ctx, "API Call", 3)
	go performTask(ctx, "Calculation", 1)
}
