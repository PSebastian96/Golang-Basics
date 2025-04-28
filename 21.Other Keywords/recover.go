package main

import "fmt"

/*
`recover` is used to regain control of a panicking goroutine.
It only works when called inside a deferred function.
*/

func safeDivide(a, b int) (result int) {
	defer func() {
		// Recover from panic and prevent the program from crashing
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			result = 0 // Provide a default value on panic
		}
	}()
	// Will panic if b is zero
	return a / b
}

func main() {
	// Example 1: Using recover to handle panic
	fmt.Println("Safe division result:", safeDivide(10, 0))
	fmt.Println("Program continues executing after recovery.")
}
