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
	fmt.Println(f)

	// iota is a special identifier that can be used in a const declaration to simplify definitions of incrementing numbers

	const (
		sunday int = iota
		monday
		tuesday
		wednesday
		thursday
		friday
		saturday
	)
	fmt.Println(sunday, monday, tuesday, wednesday, thursday, friday, saturday)

	// iota can be used to create a set of related constants
	const (
		_  = iota // skip the first iota value
		KB = 1 << (10 * iota)
		MB
		GB
		TB
	)
	fmt.Println(KB, MB, GB, TB)

}
