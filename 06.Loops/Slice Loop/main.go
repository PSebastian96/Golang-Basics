package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}

	fmt.Println("Loop through the index and values")
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	fmt.Println("\nLoop and print only the index")
	for v := range numbers {
		fmt.Printf("Index: %d\n", v)
	}

	fmt.Println("\nLoop and print only the value")
	for _, v := range numbers {
		fmt.Printf("Value: %d\n", v)
	}
}
