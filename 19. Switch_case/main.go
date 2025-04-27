package main

import "fmt"

/*
	Syntax for Pattern Matching (switch/case statement)
	switch expression {
	case value1:
		// code block
	case value2:
		// another block
	default:
		// default block if none match
	}
*/

func main() {
	day := "Wednesday"

	fmt.Printf(" > Full Switch Statement: \n")
	switch day {
	case "Monday":
		fmt.Println("Start of the work week")
	case "Tuesday":
		fmt.Println("Second day of the week")
	case "Friday":
		fmt.Println("Last working day!")
	// group multiple values in one case:
	case "Saturday", "Sunday":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Some other day")
	}

	fmt.Printf("\n > Switch Without an Expression: \n")
	age := 25

	switch {
	case age < 18:
		fmt.Println("Minor")
	case age >= 18 && age <= 60:
		fmt.Println("Adult")
	default:
		fmt.Println("Senior")
	}

	fmt.Printf("\n > Fallthrough Statement: \n")
	num := 1
	switch num {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two") // Will also print even if not matched
	}

	fmt.Printf("\n > Type Switch (Interface Handling): \n")
	var val interface{} = "hello"

	switch v := val.(type) {
	case int:
		fmt.Println("It's an int:", v)
	case string:
		fmt.Println("It's a string:", v)
	default:
		fmt.Println("Unknown type")
	}

}

/*
	SUMMARY:

	Feature 	| Description
	switch 		| Evaluates expression and selects case
	case 		| Defines possible matching values
	default 	| Runs if no case matches
	fallthrough | Forces the next case to execute
	switch {} 	| Condition-based switch (like if-else)
	type switch | Used with interfaces to get types

*/
