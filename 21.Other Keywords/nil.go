package main

import "fmt"

/*
The `nil` keyword in Go is used to represent:
- A zero value for pointers
- A zero value for channels, functions, interfaces, maps, and slices
*/

func main() {
	// 1. Using nil with pointers
	var ptr *int // nil pointer
	if ptr == nil {
		fmt.Println("Pointer is nil.")
	}

	// 2. Using nil with slices
	var slice []int // nil slice
	if slice == nil {
		fmt.Println("Slice is nil.")
	}

	// 3. Using nil with maps
	var m map[string]int // nil map
	if m == nil {
		fmt.Println("Map is nil.")
	}

	// 4. Using nil with channels
	var ch chan int // nil channel
	if ch == nil {
		fmt.Println("Channel is nil.")
	}
}
