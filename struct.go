package main

import "fmt"

func struct_() {
	// struct is a user defined type that groups together variables of different types

	var emp1 Employee
	fmt.Println(emp1) // { 0 }

	emp1.name = "John"
	emp1.age = 30
	fmt.Println(emp1)

	emp2 := Employee{name: "Harry", age: 25}
	fmt.Println(emp2)

	// accessing struct fields
	fmt.Println(emp2.name)

	// modifying struct fields
	emp2.name = "Harry Potter"

	// passing struct to function
	changeName(emp1, "John Doe")
	fmt.Println(emp1)

	changeAge(&emp1, 35)
	fmt.Println(emp1)

	emp1.SetAddress(Adress{street: "123 Main St", city: "Anytown", state: "CA"})
	fmt.Println(emp1.EmployeeInfo())
}

type Adress struct {
	street string
	city   string
	state  string
}

// embedded struct
type Employee struct {
	name    string
	age     int
	address Adress
}

func changeName(emp Employee, name string) {
	emp.name = name
	fmt.Println("inside function", emp)
}

func changeAge(emp *Employee, age int) {
	emp.age = age
	fmt.Println("inside function", *emp)
}

// method with receiver
func (emp Employee) EmployeeInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d, Address: %s, %s, %s", emp.name, emp.age, emp.address.street, emp.address.city, emp.address.state)
}

// method with pointer receiver
func (emp *Employee) SetAddress(address Adress) {
	emp.address = address
}
