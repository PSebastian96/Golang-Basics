package main

import "fmt"

/*
A struct (short for structure) is a user-defined type that allows you to group variables of different types
into a single unit.

Think of it as a blueprint for creating complex data models — like a record, object, or class without methods.
*/

/*
	Feature 		| Explanation
	Value Type 		| Structs are value types (copied on assign)
	Pointers 		| Use *Struct to modify in-place
	Exported Fields | Capitalized field names are exported (visible outside the package)
*/

func main() {

	// 1. Defined Struct
	type Book struct {
		Title   string
		Author  string
		Subject string
		BookID  int
		Read    bool
	}

	book1 := Book{
		Title:   "Shawshank Redemption",
		Author:  "Stephen King",
		Subject: "Novel about prisoners",
		BookID:  123,
		Read:    true,
	}

	fmt.Printf("Book Struct:\n")
	fmt.Println("Title:", book1.Title)
	fmt.Println("Author:", book1.Author)
	fmt.Println("Subject:", book1.Subject)
	fmt.Println("Book ID:", book1.BookID)
	fmt.Println("Read already ?:", book1.Read)

	fmt.Printf("\n")

	// 2. Anonymous struct
	person := struct {
		Name string
		Age  int
	}{
		Name: "Dave",
		Age:  40,
	}

	fmt.Printf("Anonymous Struct:\n")
	fmt.Println(person)

	// 3. Nested Struct
	type Address struct {
		City, State string
	}

	type Employee struct {
		Name    string
		Address Address
	}

	e := Employee{
		Name:    "Eva",
		Address: Address{"Austin", "Texas"},
	}

	fmt.Printf("\nNested struct:\n")
	fmt.Println("City: ", e.Address.City) // Austin
	fmt.Println("Name: ", e.Name)

}
