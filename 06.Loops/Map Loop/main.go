package main

import "fmt"

func main() {
	myMap := map[string]int{"apple": 5, "banana": 3, "cherry": 7}

	fmt.Println("Loop through index/key and value pair")
	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	fmt.Println("\nLoop and print only the Value (numbers)")
	for _, v := range myMap {
		fmt.Printf("Value: %d\n", v)
	}

	fmt.Println("\nLoop through the index (strings)")
	for i := range myMap {
		fmt.Printf("Index: %s\n", i)
	}
}
