// Functions -

package main

import "fmt"

// Functions are reusable blocks of code. They can take parameters and return results.

/*
	Function syntax:

 	func functionname(parameters parameter-datatype) returnType {
	logic of function
 	}

	Example (adds two numbers and returns an integer):
	func firstFunction(a,b int) int {
		c := a + b
		return c
	}
*/

/*
	Importing and Exporting Functions:
	1. Keep All Files in the Same Package
	2. Use Capitalized Function Names for Export
	3. Call Them in main.go
*/

/*
	If organizing files into subfolders:

	1. A different package name inside the subfolder
	2. An import path like import "projectname/mypackage"
	3. Pproper exported (Capitalized) functions
*/

func sayCapital(s string) {
	fmt.Printf("%s, is the Capital \n", s)
}
func sayCountry(s string) {
	fmt.Printf("%s, is the Country  \n", s)
}

// Finding the maximum value between two numbers
func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}

// Function to swap two strings
func swap(x, y string) (string, string) {
	return y, x
}

// Function Arguments
// Call by value
func increment(number int) {
	number++
	fmt.Println("Inside increment: ", number)
}

// Call by reference
func modify(slice []int) {
	slice[0] = 999
	fmt.Println("Inside modify: ", slice)
}

func main() {
	sayCountry("USA")
	sayCapital("Washington")
	a, b := 100, 200
	result := max(a, b)
	fmt.Printf("Max value is : %d\n", result)
	firstName, LastName := swap("John", "Doe")
	fmt.Printf("Swapped names: %s %s\n", firstName, LastName)

	// Calling increment function by value
	x := 10
	increment(x)
	fmt.Println("In main after increment:", x)

	// Calling modify function by reference
	mySlice := []int{1000, 2000, 3000}
	modify(mySlice)
	fmt.Println("In main after modify:", mySlice)
}
