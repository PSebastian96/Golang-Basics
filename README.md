# GOLANG BASICS - SYNTAX AND CODE EXAMPLES
<h1> 💻 Golang Basics - Golang CLI and Reserved Keywords</h1>


<h3>Golang CLI commands:</h3>

<ul>
  <li>go mod init filename.go -> create a module file (dependency list). </li>
  <li>go mod tidy -> It adds any missing modules necessary to build the current module's
					packages and dependencies, and it removes unused modules that
					don't provide any relevant packages. It also adds any missing entries
					to go.sum and removes any unnecessary ones. </li>
  <li> go get gorm.io/gorm -> downloads the library/dependency. </li>
  <li> go build -o filename.exe -> creates an executable file. </li>
  <li> go run filename.go OR go run filename.exe -> will run the code. </li>
  <li> go -> lists commands and explanation. </li>
  <li> go help somecommand -> more info about that command. </li>
  <li>go help topic -> more info about a topic. </li>
</ul>

<h3>Golang Reserved Keywords:</h3>

Summary Table of Reserved Keywords:

| **Category**          | **Keywords**                                                                                                                                      |
|-----------------------|---------------------------------------------------------------------------------------------------------------------------------------------------|
| **Control Flow**      | `break`, `case`, `continue`, `default`, `fallthrough`, `for`, `go`, `goto`, `if`, `else`, `select`, `switch`                                      |
| **Data Types**        | `bool`, `byte`, `complex64`, `complex128`, `float32`, `float64`, `int`, `int8`, `int16`, `int32`, `int64`, `string`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr` |
| **Declarations**      | `const`, `defer`, `func`,`map`, `var`                                                                             |
| **Concurrency**       | `chan`, `go`, `select`                                                                                                                            |
| **Types**             | `interface`, `struct`, `type`                                                                                                                     |
| **Other**             | `nil`, `panic`, `recover`, `import`, `package`

<h1>📜 Golang Formatting, Padding, Escaping and Time Format Cheatsheet</h1>

|Category | Character | Meaning | Example|
|---------|-----------|---------|--------|
|Format Specifiers | %d | Decimal integer | fmt.Printf("%d", 123) → 123
| | %s | String | fmt.Printf("%s", "hello") → hello
| | %f | Float (decimal) | fmt.Printf("%.2f", 3.1415) → 3.14
| | %t | Boolean | fmt.Printf("%t", true) → true
| | %v | Default format (any type) | fmt.Printf("%v", myVar)
| | %+v | Struct with field names | fmt.Printf("%+v", myStruct)
| | %#v | Go-syntax format | fmt.Printf("%#v", myStruct)
| | %T | Type of a variable | fmt.Printf("%T", myVar)
| | %p | Pointer address | fmt.Printf("%p", &myVar)
| | %% | Literal percent sign | fmt.Printf("%%") → %
|Aligning & Padding | %5d | Right-align with spaces | fmt.Printf("%5d", 42) → "   42"
| | %-5d | Left-align | fmt.Printf("%-5d", 42) → "42   "
| | %05d | Zero padding | fmt.Printf("%05d", 42) → "00042"
| | %8s | Right-align string | fmt.Printf("%8s", "go") → "      go"
|Escape Sequences | \n | Newline | "Hello\nWorld"
| | \t | Tab | "Hello\tWorld"
| | \\ | Backslash | "C:\\path\\file.txt"
| | \" | Double quote inside string | "She said: \"Hi\""
| | \r | Carriage return | Overwrites line (rare)
|Time Formatting | Mon Jan 2 15:04:05 MST 2006 | Base format | time.Now().Format("2006-01-02 15:04:05")
| | 2006 | Year | Format("2006") → 2025
| | 01 | Month (zero-padded) | Format("01") → 04
| | 2 | Day (no padding) | Format("2") → 9
| | 02 | Day (zero-padded) | Format("02") → 09
| | 15 | Hour (24h) | Format("15") → 23
| | 03 | Hour (12h) | Format("03") → 11
| | 04 | Minute | Format("04") → 07
| | 05 | Second | Format("05") → 45
| | PM | AM/PM | Format("PM") → PM
| | MST | Timezone | Format("MST") → UTC
|Custom Date Examples | time.Now().Format("02-Jan-2006") | 09-Apr-2025 | 
| | time.Now().Format("2006/01/02 15:04") | 2025/04/09 22:45 | 
| | time.Now().Format("Mon, 02 Jan 2006") | Wed, 09 Apr 2025 | 


<h1>🎯 Go (Golang) CLI Workflow for Projects</h1>

<h3>1. 📁 Create project directory</h3>
<p>✅ Always work inside a clean directory for each project.</p>

<ul>
<li>mkdir my-awesome-project</li>
<li>cd my-awesome-project</li>
</ul>

<h3>2. 🧩 Initialize a Go Module</h3>
<p>✅ This creates a go.mod file that manages all dependencies.</p>
<ul>
<li>`go mod init filename.go`</li>
<br>
<p>OR</p>
<li>`go mod init my-awesome-project` (directory name)</li>
<br>
<p>OR</p>
<li>`go mod init github.com/yourusername/my-awesome-project`</li>
</ul>

<h3>3. 📄 Create main.go</h3>

<p>✅ This is the  app’s entry point.</p>
<ul>
	<li>`touch main.go`</li>
</ul>

<h3>4. 🛠️ Create directories for clean architecture</h3>
<p>✅ Create directories quickly:</p>

![image](https://github.com/user-attachments/assets/e854db07-b5a3-49cc-9ca2-0d07282a0a19)

<ul>
<li>`mkdir controllers models routes services config`</li>
</ul>

<h3>5. 📚 Adding files</h3>
<p>Example to create a new Go file inside models:</p>
<ul>
<li>touch models/user.go
</li>
</ul>

<h3>6. 🔥 Installing libraries</h3>
<p>✅ go get automatically updates your go.mod and go.sum.</p>

![image](https://github.com/user-attachments/assets/561e94a7-ead4-4997-a330-93e1a2261598)


<h3>7. 🧪 Run your app</h3>
<p>✅ This runs the application.</p>
<ul>
<li>go run main.go</li>
</ul>

<h3>8. 🚀 Build your app (for production)</h3>
<p>This creates an executable file like `my-awesome-project.exe` (Windows) or `my-awesome-project` (Linux/Mac).</p>
