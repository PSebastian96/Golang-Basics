// Go Recursion - Notes and Examples
package main

import "fmt"

/*
Notes:
- Recursion is when a function calls itself to solve smaller subproblems.
- Every recursive function must have a base case to stop the recursion.
- Common examples: factorial, fibonacci, sum, tree/graph traversal, etc.
- Too deep recursion can cause stack overflow (Go has limited stack size).
- For performance, consider memoization or iteration (especially for fibonacci).
- Use recursion when the problem is naturally recursive or tree-like.
*/

// Factorial using recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// Fibonacci using recursion
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Sum of numbers recursively
func recursiveSum(n int) int {
	if n == 0 {
		return 0
	}
	return n + recursiveSum(n-1)
}

func main() {
	fmt.Println("--- Recursion Examples ---")
	fmt.Println("Factorial of 5:", factorial(5))
	fmt.Println("Fibonacci of 7:", fibonacci(7))
	fmt.Println("Sum of first 5 numbers:", recursiveSum(5))
}
