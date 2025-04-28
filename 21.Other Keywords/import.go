package main

import (
	"fmt"   // Importing the "fmt" package for formatted I/O
	"math"  // Importing the "math" package for mathematical operations
)

/*
`import` is used to include external packages or libraries into your program.
You can import one or more packages in one statement.
*/

func main() {
	// 1. Using the fmt package to print to the console
	fmt.Println("Hello, Go!")

	// 2. Using the math package to calculate the square root
	fmt.Println("Square root of 25:", math.Sqrt(25))
}
