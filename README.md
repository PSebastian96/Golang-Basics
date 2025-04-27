<h1>Golang Basics - Golang syntax and code examples</h1>

<h3>Golang CLI commands:</h3>

<ul>
  <li>go mod init filename.go -> create a module file (dependency list)</li>
  <li>go mod tidy -> It adds any missing modules necessary to build the current module's
					packages and dependencies, and it removes unused modules that
					don't provide any relevant packages. It also adds any missing entries
					to go.sum and removes any unnecessary ones.</li>
  <li> go build -o filename.exe -> creates an executable file </li>
  <li> go run filename.go OR go run filename.exe -> will run the code </li>
  <li> go -> lists commands and explanation </li>
  <li> go help somecommand -> more info about that command</li>
  <li>go help topic -> more info about a topic</li>
</ul>

<h3>Golang Reserved Keywords:</h3>

Summary Table of Reserved Keywords:

| **Category**          | **Keywords**                                                                                                                                      |
|-----------------------|---------------------------------------------------------------------------------------------------------------------------------------------------|
| **Control Flow**      | `break`, `case`, `continue`, `default`, `fallthrough`, `for`, `go`, `goto`, `if`, `else`, `select`, `switch`                                      |
| **Data Types**        | `bool`, `byte`, `complex64`, `complex128`, `float32`, `float64`, `int`, `int8`, `int16`, `int32`, `int64`, `string`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr` |
| **Declarations**      | `const`, `defer`, `func`, `interface`, `map`, `struct`, `type`, `var`                                                                             |
| **Concurrency**       | `chan`, `go`, `select`                                                                                                                            |
| **Types**             | `interface`, `struct`, `type`                                                                                                                     |
| **Other**             | `nil`, `panic`, `recover`, `import`, `package`
