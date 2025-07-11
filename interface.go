package main

import "fmt"

func interface_() {

	rect := Rectangle{Width: 10, Height: 5}
	area, perimeter := CalculateArea(rect)

	circle := Circle{Radius: 5}
	area2, perimeter2 := CalculateArea(circle)

	fmt.Printf("Area and perimeter of rectangle is with width %v and height %v is %v and %v", rect.Width, rect.Height, area, perimeter)
	fmt.Println()
	fmt.Printf("Area of circle is with radius %v is %v and %v", circle.Radius, area2, perimeter2)

	// empty interface
	describe(10)
	describe("10")
	describe(true)

	var myAny any // any is alias for interface{}
	myAny = 10
	fmt.Printf("(%v, %T)\n", myAny, myAny)
	myAny = "10"
	fmt.Printf("(%v, %T)\n", myAny, myAny)

	// type assertion
	var i interface{} = "hello"

	if s, ok := i.(string); ok {
		fmt.Println("s is a string", s)
	} else {
		fmt.Println("s is not a string")
	}

	var new_rect Shape = Rectangle{Width: 10, Height: 5}

	if r, ok := new_rect.(Rectangle); ok {
		fmt.Println("r is a Rectangle", r)
	}

	// type switch
	switch v := new_rect.(type) {
	case Circle:
		fmt.Println("r is a Circle", v)
	case Rectangle:
		fmt.Println("r is a Rectangle", v)
	default:
		fmt.Println("r is unknown type")
	}

	// pointer vs value receiver
	person1 := Person{name: "John", age: 30}
	person1.SetName("Harry")
	fmt.Println(person1.name, person1.age)
	person1.SetAge(25)
	fmt.Println(person1.name, person1.age)

}

type Shape interface {
	Area() float64
}
type Measurable interface {
	Perimeter() float64
}

type Geometry interface {
	Shape
	Measurable
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func CalculateArea(g Geometry) (float64, float64) {
	return g.Area(), g.Perimeter()
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type Person struct {
	name string
	age  int
}

type MyInterface interface {
	SetName() string
	SetAge() int
}

func (p Person) SetName(name string) string {
	p.name = name
	return p.name
}

func (p *Person) SetAge(age int) int {
	p.age = age
	return p.age
}
