package main

import "fmt"

// while loop example

func main() {
	count := 0

	// while count is less than 5
	for count < 5 {
		fmt.Println("Count is : ", count)
		count++
	}

	fmt.Println("Done!")
}
