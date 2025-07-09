package main

import "fmt"

func pointer() {
	a := 10
	var p *int = &a
	fmt.Println("address of a", p)
	fmt.Println("value of a", *p)

	// pointer with function
	b := 20
	modifyByVal(b)
	fmt.Println("after modifyByVal", b)
	modifyByRef(&b)
	fmt.Println("after modifyByRef", b)

	// new vs make
	// new : allocate memory for a variable and return a pointer to it
	// make : allocate memory for a variable and return a reference to it
	// new : used for primitive types
	// make : used for slices, maps, channels
	// new : return a pointer to a zero value
	// make : return a reference to a initialized value
	c := new(int)
	fmt.Printf("new(int): Value: %d, Type: %T, Address: %p\n", *c, c, c)
	*c = 10
	fmt.Printf("new(int) after assignment: Value: %d, Type: %T, Address: %p\n", *c, c, c)

}

func modifyByVal(a int) {
	a = 30
	println("inside function", a)
}

func modifyByRef(a *int) {
	*a = 30
	println("inside function", *a)
}
