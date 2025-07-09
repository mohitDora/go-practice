package main

import (
	"fmt"
)

func variable() {

	// declare variable
	var a int
	var name string
	var isTrue bool

	// assign value
	a = 10
	name = "John"
  isTrue = true

  // declare and assign value
  var city string = "New York"

  /* if value is not assigned, it will be set to default value 
  int, float : 0
  string : ""
  bool : false
  */

	fmt.Println(a, name, isTrue, city)

  // type inference
  var b = 20
  //b=3.14 // error: cannot use 3.14 (untyped float constant) as int value in assignment
  fmt.Printf("Type of b is %T\n", b)

  // short declaration
  c := 30
  fmt.Println(c)

  /* Derived types 
  - Pointer
  - Array
  - Struct
  - Channel
  - Function
  - Interface
  - Map
  - Slice
  */

  // type conversion
  var d float64 = 40.5
  var e int = int(d)

  fmt.Println(d, e)
  
}
