// Go Maps - Notes and Examples
package main

import "fmt"

/*
Notes:
- Maps are key-value pairs, similar to dictionaries in Python.
- Keys must be unique and comparable (no slices, maps, or functions).
- Use `make()` or map literals to create maps.
- Accessing a nonexistent key returns the zero value for the type.
- Use the `comma ok` idiom to check if a key exists.
- `delete(map, key)` removes a key-value pair.
- Maps are reference types, passed by reference.
- Map iteration order is not guaranteed (it's randomized).
*/

// Syntax: map[key_data_type]value_data_type{key1: value1, key2: value2}

func main() {
	fmt.Println("--- Creating Maps ---")

	// 1. Declare and initialize with literal
	m1 := map[string]int{"Alice": 25, "Bob": 30}
	fmt.Println("m1:", m1)

	// 2. Using make()
	m2 := make(map[string]string)
	m2["country"] = "Canada"
	m2["language"] = "English"
	fmt.Println("m2:", m2)

	fmt.Println("--- Accessing Values ---")
	fmt.Println("Alice's age:", m1["Alice"])

	// Check if key exists
	val, ok := m1["Eve"]
	if ok {
		fmt.Println("Eve found with age:", val)
	} else {
		fmt.Println("Eve not found")
	}

	fmt.Println("--- Iterating Over Map ---")
	for key, value := range m1 {
		fmt.Printf("%s: %d\n", key, value)
	}

	fmt.Println("--- Deleting from Map ---")
	delete(m1, "Bob")
	fmt.Println("After deletion:", m1)

	fmt.Println("--- Length of Map ---")
	fmt.Println("len(m1):", len(m1))
}
