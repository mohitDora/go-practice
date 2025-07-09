package main

import "fmt"

func function() {
	fmt.Println(add(2, 2))
	fmt.Println(add2(2, 10))
	fmt.Println(add3(2, 10))
	fmt.Println(add4(1, 2, 3, 4, 5))

	counter1 := counter()
	counter2 := counter()
	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter2())
	fmt.Println(counter2())

	func() {
		fmt.Println("IIFE: Immediately Invoked Function Expression")
	}()

}

// function with return value
func add(a int, b int) int {
	return a + b
}

// function with multiple return values
func add2(a, b int) (int, bool) {
	sum := a + b
	if sum > 10 {
		return sum, true
	}
	return sum, false
}

// function with named return values
func add3(a, b int) (sum int) {
	sum = a + b
	return
}

// function with variadic parameters
func add4(a ...int) (sum int) {
	for _, v := range a {
		sum += v
	}
	return
}

// function with anonymous function
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
