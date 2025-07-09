package main

import (
	"errors"
	"fmt"
)

func error_handling() {

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result2, err2 := divide2(10, 0)
	if err2 != nil {
		fmt.Println("Error:", err2)
		var divisionError *DivisionError
		if errors.As(err2, &divisionError) {
			fmt.Println("Division error", divisionError.Dividend, divisionError.Divisor)
		}
	} else {
		fmt.Println("Result:", result2)
	}

	err3 := fileOperation("non_existent.txt")
	if err3 != nil {
		fmt.Println("Error:", err3)
		if errors.Is(err3, ErrFileNotFound) {
			fmt.Println("File not found")
		}
	}

	// panic and recover
	// panic : stop the normal execution of the current function and begin unwinding the stack
	// recover : stop the unwinding of the stack and return the value passed to panic and only works inside deferred functions

	mightPanic(2)
	mightPanic(0)

}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// custom error struct
type DivisionError struct {
	Dividend int
	Divisor  int
	Message  string
}

// implement error interface
func (e *DivisionError) Error() string {
	return fmt.Sprintf("%s: %d / %d", e.Message, e.Dividend, e.Divisor)
}

// custom error function
func divide2(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivisionError{a, b, "division by zero"}
	}
	return a / b, nil
}

// error wrapping
var ErrFileNotFound = errors.New("file not found")
var ErrFilePermission = errors.New("permission denied")

// error wrapping function
func fileOperation(fileName string) error {
	if fileName == "non_existent.txt" {
		return ErrFileNotFound
	}
	if fileName == "protected.txt" {
		return ErrFilePermission
	}
	return nil
}

// panic and recover
func mightPanic(num int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	res := 10 / num
	fmt.Println(res)
}
