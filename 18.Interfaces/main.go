package main

import (
	"fmt"
	"math"
)

/*
	--- Interfaces Notes ---
	- Interfaces define a set of method signatures that a type must implement.
	- They provide a way to achieve polymorphism (behavior over different types).
	- Any type that implements all methods of an interface is said to satisfy that interface.
	- Interfaces are useful for abstraction, testing, and designing extensible systems.
	- You can use type assertions or type switches to extract underlying types.
*/

// Interface -> similat to struct, it holds functions instead of variables

/* 	Syntax of an INTERFACE:
type thisisaninterface interface {
	// MethodName1 ReturnType
	// MethodName2 ReturnType
	// MethodName3 ReturnType
}

Syntax of a STRUCT:
type thisisastruct struct {
	// variables
}
*/

// Define an interface
type Shape interface {
	area() float64
}

// Define a Circle struct
type Circle struct {
	radius float64
}

// Define a Rectangle struct
type Rectangle struct {
	width, height float64
}

// Implement area method for Circle (Struct is used as paramater, function name is used from Interface)
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Implement area method for Rectangle (Struct is used as paramater, function name is used from Interface)
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// Function to get area of any shape!
func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	fmt.Println("Interfaces")

	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
}

/* Notes:

The function area() has been declared for both Circle and Rectangle structs,
because the variables circle and rectangle have been declared as their respected structs,
therefore golang automatically in runtime checks and compares which functions to use in order to compute the result.

If the area() function is removed ( example:func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
} )

it will return an error for circle.

*/
