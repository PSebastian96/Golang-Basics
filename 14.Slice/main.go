// Go Slices - Notes and Examples
package main

import "fmt"

/*
Notes:
- Slices are more powerful than arrays. They are dynamic and reference-based.
- The underlying array is shared between slices, so changes reflect unless a copy is made.
- Length is number of elements. Capacity is from start to end of underlying array.
- Use `make` to initialize slices with predefined length/capacity.
- Use `append` to add elements (may reallocate underlying array).
- Always be aware of shared memory when slicing.
*/

func main() {
	fmt.Println("--- Creating Slices ---")

	// 1. Declare and Initialize
	slice1 := []int{1, 2, 3, 4, 5} // literal slice
	fmt.Println("slice1:", slice1)

	// 2. Using make()
	slice2 := make([]int, 5) // length 5, default values (0)
	fmt.Println("slice2:", slice2)

	// 3. Slicing an Array
	arr := [5]int{10, 20, 30, 40, 50}
	slice3 := arr[1:4] // elements at index 1, 2, 3
	fmt.Println("slice3 (from array):", slice3)

	fmt.Println("--- Accessing and Modifying ---")
	slice1[0] = 100
	fmt.Println("Updated slice1:", slice1)

	fmt.Println("--- Appending Elements ---")
	slice1 = append(slice1, 6, 7)
	fmt.Println("Appended slice1:", slice1)

	fmt.Println("--- Length and Capacity ---")
	fmt.Println("Length:", len(slice1), "Capacity:", cap(slice1))

	fmt.Println("--- Iterating Over Slices ---")
	for i, v := range slice1 {
		fmt.Printf("Index %d = %d\n", i, v)
	}

	fmt.Println("--- Copying Slices ---")
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println("Copied dst:", dst)

	fmt.Println("--- Slice Internals ---")
	s1 := []int{1, 2, 3}
	s2 := s1[:2]
	s2[0] = 99
	fmt.Println("s1:", s1) // s1 is affected!
	fmt.Println("s2:", s2)

	// Subslicing
	numbers2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	// inclunsive of 1st index
	subSlice1 := numbers2[1:4]
	// inclusive of 0th index, start of slice
	subSlice2 := numbers2[:3]
	// inclusive of the given index, 4th index till the end
	subSlice3 := numbers2[4:]
	fmt.Println("subSlice1:", subSlice1)
	fmt.Println("subSlice2:", subSlice2)
	fmt.Println("subSlice3:", subSlice3)
}

/*

cap() -> Returns the number of elements between the start of the slice and the end of the underlying array.
It’s not how many are left, but how much total space is available from where your slice starts.

Why is it useful?
 To understand how much room you have left to grow (before append() triggers a new array allocation).
 When slicing or indexing, to avoid out-of-bounds errors.
 To write optimized code, especially in loops or performance-critical paths.

 cap() is super useful for checking how much buffer space you've got left after slicing or appending.
 Especially helpful when building things like custom buffer pools or performance-tuned code.

*/
