package main

import "fmt"

// Arrays are fixed-size sequences of elements with the same type.
// Zero Indexing: The first element is accessed with index 0.

/*
	Arrays:

	Fixed size
	Value type
	Not often used directly for flexible data

	Slices:

	Dynamic, flexible
	Preferred in most real-world apps

	Mistake:					| Fix:
	Using out-of-range index	| Always check len(arr)
	Thinking arrays are dynamic	| Use slices instead
	Expecting mutation via func	| Pass a pointer or use slice
*/

// var arr [3]int -> declares an array of 3 integers

// arr := [4]string{"A", "B", "C", "D"} -> length is fixed at 4

func main() {

	// fixed array
	a := [3]int{1, 2, 3}
	b := a
	b[0] = 99
	fmt.Printf("Fixed size array: \n")
	fmt.Println(a) // still [1 2 3], not changed
	fmt.Println(b) // will return [99 2 3] changed array

	// Use [...] to Auto-Detect Size
	primes := [...]int{2, 3, 5, 7, 11}
	fmt.Printf("\n")
	fmt.Printf("[...] auto detect the size of array:\n")
	fmt.Println(len(primes)) // 5
	fmt.Printf("Primes array values: %d", primes)

}
