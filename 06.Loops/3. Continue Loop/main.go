package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // skip even numbers
		}
		fmt.Println(i)
	}
}
