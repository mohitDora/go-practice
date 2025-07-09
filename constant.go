package main

import "fmt"

func constant() {
	const pi = 3.14
	fmt.Println(pi)

	// multiple constant declaration
	const (
		a = 1
		b = 2
		c = 3
	)
	fmt.Println(a, b, c)

	const untypedConst = 10
	var d int = untypedConst
	var e float64 = untypedConst
	fmt.Println(d, e)

	const typedConst int = 10
	var f int = typedConst
	// var g float64 = typedConst // error: cannot use typedConst (untyped int constant) as float64 value in variable declaration
	fmt.Println(f, g)
}
