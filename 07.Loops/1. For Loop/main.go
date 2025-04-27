package main

import "fmt"

// Loops or iteration means to go through items or repeat something as long as it has achieved its goal
// for loop example

func main() {
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even number")
		} else {
			fmt.Println(i, "is odd number")
		}
	}

	fmt.Println("For loop finished!")
}

// 'for' keyword is used to iterate/loop through items and data types.
// loops can NOT be declared outside a funcion
// struct data types don't support direct iteration

/*
	Keyword - Description
	range - loop through data types
	break - Stops the loop
	continue - Skips to the next iteration
	return - Exits the loop and the function
	goto - Jumps to a labeled statement (rarely used)
*/
