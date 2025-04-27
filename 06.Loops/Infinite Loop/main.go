package main

import "fmt"

// infinite loop example

func main() {
	for {
		fmt.Println("Running...")
		return // exits the loop and the function
	}
}

/*
	for {
		fmt.Println("This runs forever")
	}

	This would run forever, until the computer shuts down
	In order to avoid forever loops, use 'break' keyword or
	conditional when to stop.

	Keyword - Description
	break - Stops the loop
	continue - Skips to the next iteration
	return - Exits the loop and the function
	goto - Jumps to a labeled statement (rarely used)
*/
