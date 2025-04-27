package main

import "fmt"

// do while example

func main() {
	count := 0

	for {
		fmt.Println("Count is: ", count)
		count++

		if count >= 5 {
			break
		}
	}

	fmt.Println("Do While Loop finished!")
}
