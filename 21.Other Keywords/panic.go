package main

import "fmt"

/*
`panic` is used to cause the program to terminate immediately and begin unwinding the stack.
It is used to indicate a critical error where the program cannot recover.
*/

func divide(a, b int) int {
	if b == 0 {
		panic("division by zero") // Panics if b is zero
	}
	return a / b
}

func main() {
	// Example 1: Panic when dividing by zero
	fmt.Println("Starting the program...")
	fmt.Println("Result:", divide(10, 0)) // This will panic
}
