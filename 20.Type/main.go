package main

import "fmt"

/*
In Go, the `type` keyword is used to:
- Define new types (like structs, slices, or even your own simple types)
- Create type aliases
- Build complex models

It's very powerful for abstraction, readability, and enforcing stronger typing.
*/

/*
	Feature 		| Explanation
	Custom Types 	| Create your own named types
	Aliases 		| Create an alias for an existing type
	Structs 		| Group fields together under one type
*/

// 1. Custom simple types
type Age int
type Name string

// 2. Type alias
type Score = int // `Score` is now just another name for int

// 3. Struct type
type Book struct {
	Title   string
	Author  string
	Pages   int
	IsRead  bool
}

// 4. Struct with nested struct
type Address struct {
	City, Country string
}

type User struct {
	Name    Name
	Age     Age
	Address Address
}

func main() {
	// Using custom types
	var myAge Age = 30
	var myName Name = "John Doe"

	fmt.Printf("Custom types:\n")
	fmt.Println("Name:", myName)
	fmt.Println("Age:", myAge)

	fmt.Printf("\n")

	// Using alias type
	var mathScore Score = 95
	fmt.Printf("Alias type:\n")
	fmt.Println("Math Score:", mathScore)

	fmt.Printf("\n")

	// Using struct
	book := Book{
		Title:  "The Go Programming Language",
		Author: "Alan A. A. Donovan",
		Pages:  380,
		IsRead: false,
	}

	fmt.Printf("Struct example:\n")
	fmt.Println("Book Title:", book.Title)
	fmt.Println("Author:", book.Author)
	fmt.Println("Pages:", book.Pages)
	fmt.Println("Have you read it?", book.IsRead)

	fmt.Printf("\n")

	// Using nested struct
	user := User{
		Name: "Alice",
		Age:  28,
		Address: Address{
			City:    "Berlin",
			Country: "Germany",
		},
	}

	fmt.Printf("Nested Struct example:\n")
	fmt.Println("User Name:", user.Name)
	fmt.Println("City:", user.Address.City)
	fmt.Println("Country:", user.Address.Country)
}
