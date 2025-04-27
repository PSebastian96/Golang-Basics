package main

import "fmt"

func main() {
	var i int = 0

loopStart:
	if i >= 5 {
		fmt.Println("Done looping!")
		return
	}
	fmt.Printf("i = %d\n", i)
	i++
	goto loopStart
}

/*
	- loopStart: is a label (like a named anchor).
	- goto loopStart tells Go to jump back to that label.
	- It keeps jumping back until the condition i >= 5 becomes true.

	Using goto is not idiomatic Go in most cases — it's generally discouraged unless you're:
	- Exiting deeply nested loops
	- Handling error cleanup in low-level code
	- Writing state-machine-style logic
	- Idiomatic Go prefers structured loops like for, break, and continue.
*/
