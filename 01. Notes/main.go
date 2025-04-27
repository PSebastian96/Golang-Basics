package main

// GOLANG NOTES

/*
	Summary Table of Reserved Keywords:

	| **Category**          | **Keywords**                                                                                                                                      |
	|-----------------------|---------------------------------------------------------------------------------------------------------------------------------------------------|
	| **Control Flow**      | `break`, `case`, `continue`, `default`, `fallthrough`, `for`, `go`, `goto`, `if`, `else`, `select`, `switch`                                      |
	| **Data Types**        | `bool`, `byte`, `complex64`, `complex128`, `float32`, `float64`, `int`, `int8`, `int16`, `int32`, `int64`, `string`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr` |
	| **Declarations**      | `const`, `defer`, `func`, `interface`, `map`, `struct`, `type`, `var`                                                                             |
	| **Concurrency**       | `chan`, `go`, `select`                                                                                                                            |
	| **Types**             | `interface`, `struct`, `type`                                                                                                                     |
 	| **Other**             | `nil`, `panic`, `recover`, `import`, `package`

*/

/*

	Golang CLI commands:

	go mod init filename.go -> create a module file (dependency list)
	go mod tidy -> It adds any missing modules necessary to build the current module's
					packages and dependencies, and it removes unused modules that
					don't provide any relevant packages. It also adds any missing entries
					to go.sum and removes any unnecessary ones.

	go build -o filename.exe -> creates an executable file

	go run filename.go OR go run filename.exe -> will run the code

	go -> lists commands and explanation

	go help somecommand -> more info about that command

	go help topic -> more info about a topic

*/
