package main

import "fmt"

// single line comment

/*
Multi
line
comment
*/

// Data Types: Integers, floating points, booleans, complex numbers and strings

// Funcion syntax
func main() {
	fmt.Println("----- From main() Function -----")
	fmt.Println("Hello World!") // ; semi-colon can be ignored
	fmt.Println("---------\n ---------------")

	// Section1: Integers
	// Section1-A: signed integers - Can be positive and negative numbers, default value is 0
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	// Assigning values to integer variables
	i = -128
	i8 = 127
	i16 = -32768
	i32 = 2147483647
	i64 = -9223372036854775808

	// Section1-B: unsigned integers - Only positive numbers
	var u uint
	var u8 uint8
	var u16 uint16
	var u32 uint32
	var u64 uint64

	// Assigning values to unsigned integer variables
	u = 255
	u8 = 255
	u16 = 65535
	u32 = 4294967295
	u64 = 18446744073709551615

	// Printing the signed and unsigned integer variables to the console
	fmt.Println("------ Integers ------")
	fmt.Println("Signed integers:")
	fmt.Println("i:", i)
	fmt.Println("i8", i8)
	fmt.Println("i16:", i16)
	fmt.Println("i32:", i32)
	fmt.Println("i64:", i64)

	fmt.Println("Unsigned integers:")
	fmt.Println("u:", u)
	fmt.Println("u8:", u8)
	fmt.Println("u16:", u16)
	fmt.Println("u32:", u32)
	fmt.Println("u64", u64)
	fmt.Print("------------------------\n")

	// Section2: Floating Point - by default supports positive and negative numbers
	// Section2-A : Float32
	// This is used for less precise calculations
	var f32 float32 = 10.6
	// Section2-B : Float64
	var f64 float64 = 10.6
	// This is used for more precise calculations

	// Printing the floating point values
	fmt.Println("------ Floating Point Values ------")
	fmt.Println("FLOAT32:", f32)
	fmt.Println("FLOAT64:", f64)
	fmt.Print("------------------------\n")

	// Demonstrating the diff in precision between the f32 and f64
	var HP float64 = 10123456789012345
	var LP float32 = 10123456789012345
	fmt.Println("------ Difference f32 and f64 ------")
	fmt.Println("High precision float64:", HP)
	fmt.Println("Low precision float32:", LP)
	fmt.Print("------------------------\n")

	// Section3: Boolean Data Type - true or false values, default value false
	var isActive bool = true
	var isOn bool = false

	fmt.Println("------ Boolean Values ------")
	fmt.Println("Is Active:", isActive)
	fmt.Println("Is On:", isOn)
	fmt.Print("------------------------\n")

	// Section4: Complex Data Type
	var CN1 complex128 = complex(5, 10)
	var CN2 complex64 = complex(2, 7)
	// print(CN1)
	// print(CN2)
	fmt.Println("------ Complex Number Values ------")
	fmt.Println("CN1: ", CN1)
	fmt.Println("CN2: ", CN2)
	fmt.Print("------------------------\n")

	// Section5: String Data Type - usually text, can contain numbers, defualt value "" (empty string)
	var name string = "Kermet the frog!"
	fmt.Println("------ Strings ------")
	fmt.Println("My name is:", name)
	fmt.Println("----------------------")

}
