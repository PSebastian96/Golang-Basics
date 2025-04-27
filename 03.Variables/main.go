package main

import "fmt"

// Variables can be assigned in 2 ways: one line and different lines
// Variables store values with a specific type.

func main() {

	// Variables-A: Declare variables on different lines (depricated form GO ver 1.23 )
	var mango string
	mango = "Mango is a fruit!"
	var weight int
	weight = 54

	// Variables-B: Declare in one line (Encouraged from GO ver 1.23)
	var height int = 23

	// Print the variables
	fmt.Println("Mango: ", mango)
	fmt.Println("weigt: ", weight)
	fmt.Println("Height: ", height)

	// Variables-C: : ':=' used within functions to declare and assign value to the variable
	// 'age := 42' is same as 'var age = 42'

	age := 42
	fmt.Println("Age: ", age)
	city := "Cardiff"
	fmt.Println("City: ", city)

	// Variables-D : multiple declaration in one line
	var apples, oranges int = 23, 78
	fmt.Println("I have", apples, "apples and", oranges, "oranges.")

	var fruits = apples + oranges
	fmt.Println("Total fruits: ", fruits)

	var windows, mac, linux string = "Windows is ok\n", "Mac is meh\n", "Linux is GOAT!"
	fmt.Println(windows, mac, linux)

	// Static Type Declaration

	var x float64 = 20.5
	fmt.Println(x)
	fmt.Printf("X is of type:  %T\n", x)

	// Dynamic Type Declaration
	y := 89
	fmt.Println(y)
	fmt.Printf("Y is of type: %T\n", y)

	// Mixed Variable Declaration
	var a, b, c = 758.52, 8, "foobar"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a is of type: %T\n", a)
	fmt.Printf("b is of type: %T\n", b)
	fmt.Printf("c is of type: %T\n", c)
}
