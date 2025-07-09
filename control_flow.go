package main

import "fmt"

func control_flow() {
	// if else
	a := 10
	if a > 5 {
		fmt.Println("a is greater than 5")
	} else if a == 5 {
		fmt.Println("a is equal to 5")
	} else {
		fmt.Println("a is less than 5")
	}

	// switch
	user := "admin"
	switch user {
	case "admin":
		fmt.Println("user is admin")
	case "employee":
		fmt.Println("user is employee")
	default:
		fmt.Println("user is unknown")
	}

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// infinite loop
	j := 0
	for {
		fmt.Println("infinite loop")
		j++
		if j == 5 {
			break
		}
	}

	// for range
	arr := []int{1, 2, 3, 4, 5}
	for index, value := range arr {
		fmt.Println(index, value)
	}

	/* defer : defer a function call until the surrounding function returns
	This is incredibly useful for ensuring that cleanup code (like closing files, releasing locks, or closing network connections)
	*/
	exampleDefer()
}

func exampleDefer() {
	fmt.Println("Starting function")

	defer fmt.Println("Deferred message 1: This runs third")
	defer fmt.Println("Deferred message 2: This runs second")
	defer fmt.Println("Deferred message 3: This runs first")

	fmt.Println("Function body processing...")

	fmt.Println("Function body finished.")
}
