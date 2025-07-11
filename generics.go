package main

import "fmt"

func generics() {
	// generics allow us to write functions and data structures that can work with any data type
	// generics are implemented using type parameters
	// type parameters are specified using square brackets
	// type parameters can be used to specify the type of the function parameters and return type
	// type parameters can be used to specify the type of the data structure elements

	// How generics work
	// When we use a generic function or data structure, the compiler creates a new version of the function or data structure with the specified type parameters. This is called instantiation.

	fmt.Println(Greater(1, 2))
	fmt.Println(Greater(2.5, 2))
	fmt.Println(Greater("hi", "hello")) // error because string does not implement Number interface

	// generic data structure
	stack := Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack)
	if item, err := stack.Pop(); err == nil {
		fmt.Println(item)
	} else {
		fmt.Println(err)
	}
	if item, err := stack.Peek(); err == nil {
		fmt.Println(item)
	} else {
		fmt.Println(err)
	}
	fmt.Println(stack.IsEmpty())
}

type Number interface {
	int | float64
}
type String interface {
	string
}

func Greater[T Number | String](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// generic data structure
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.items) == 0 {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if len(s.items) == 0 {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
