package main

import (
	"fmt"
	"strconv"
)

/*
	--- Type Casting Notes ---
	- Type casting (type conversion) is converting a variable from one type to another.
	- Use syntax: `T(v)` where T is the target type, and v is the value.
	- Common conversions: int to float64, float64 to int, int to uint, etc.
	- For strings to numbers (or vice versa), use the `strconv` package.
	- Invalid casts (e.g. from string to int directly) require parsing and can error.
*/

func main() {
	fmt.Println("--- Type Casting Examples ---")
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Printf("int: %d, float64: %.2f, uint: %d\n", i, f, u)

	// String to int (with strconv)
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Conversion error:", err)
	} else {
		fmt.Println("Converted string to int:", num)
	}
}
