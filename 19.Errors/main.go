package main

import (
	"errors" // imported to create error messages
	"fmt"
	"math"
)

// The error interface
type error interface {
	Error() string
}

// Sqrt calculate the square root of a function
func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("negative number passed to Sqrt")
	}
	// Return the square root, and a nil error
	return math.Sqrt(value), nil
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {

	// Test the Sqrt function with a negative value
	result, err := Sqrt(-5)
	if err != nil {
		fmt.Println("Sqrt Error:", err)
	} else {
		fmt.Println("Sqrt Result:", result)
	}

	// Positive number:
	// result, err := Sqrt(5)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }

	result2, err := divide(10, 0)
	if err != nil {
		fmt.Println("Division Error:", err)
		return
	}
	fmt.Println("Division Result:", result2)
}

/*
	Concept Description:

	error 		 | Built-in interface for errors
	errors.New   | Creates a simple error
	fmt.Errorf 	 | Creates formatted or wrapped errors
	panic 		 | Stops normal flow for critical issues
	recover 	 | Used to catch panics in defer
	errors.Is/As | Compare or type-check error values

*/
