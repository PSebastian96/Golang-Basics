package main

import "fmt"

/*
A pointer is a variable that stores the memory address of another variable.
Think of it like a signpost 🪧 pointing to where the value is stored, rather than holding the value itself.
*/

// var p *int -> Declare a pointer

/*
x := 10
p := &x -> p now holds the address of x
fmt.Println(*p) -> prints 10 (the value at address p)

Storing the address: & to get the address of a variable.
Dereferencing: * to access the value at the pointer's address

*/

/*
	Key Takeaways:
	1 Pointers store memory addresses, enabling direct access to data.
	2 Use & to get an address and * to access the value at that address.
	3 Nil pointers represent uninitialized pointers.
*/

/*
	Use Case 							| Why it Matters
	Modify a variable inside a function | Pass its pointer so you can change it
	Efficiently handle large structs 	| Pass by reference instead of copying
	Shared state 						| Multiple parts of code can modify same var
*/

func main() {
	x := 42
	p := &x

	fmt.Println("x:", x)
	fmt.Println("Address of x (p):", p)
	fmt.Println("Value at p:", *p)

	*p = 99 // modify the value at that address
	fmt.Println("x after change via pointer (*p = 99):", x)

	// Nil Pointers
	var ptr *int
	fmt.Printf("The value of ptr is : %x nil pointer\n", ptr)
}
