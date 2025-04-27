// Scopes in Go.

package main

import "fmt"

// 1. Global Variables
// Global variable declaration
var g int

func main() {
	// 2. Local Variables
	// Local variable declaration
	var a, b int

	// Variable initialization
	a = 20
	b = 20
	g = a + b

	var x int = 100
	var y int = 1000

	// Displaying the values of the variables
	fmt.Printf("a = %d, b = %d, global variable g = %d\n", a, b, g)
	fmt.Printf("x = %d\n", x)
	fmt.Printf("This y local variable shadows the global variable = %d\n", y)
}
