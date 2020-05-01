package main

import "fmt"

/*
- Slices are similar to arrays, only more powerful
- Their size is not fixed and can be updated dynamically
- Indexing starts from 0
- Slicing is half-open
- However, while dynamically increasing or decreasing their size, slices use more memrory. So if you know the size of your container, arrays are a better choice. Use slices judiciously.
*/

// Declaring slices in a function
func dummyFunc(t []int) []int {
	t = append(t, 5)
	t = append(t, 4)
	return t
}

func main() {
	// Declare a slice using the built-in make function. Here 3 is the length of the slice
	s := make([]string, 3)

	fmt.Println("s empty:", s)

	// Fill the slice using indices
	s[0] = "Go"
	s[1] = "is"
	s[2] = "awesome"

	fmt.Println("s init:", s)

	// Append new values to the end of slice
	s = append(s, "let's")
	s = append(s, "go")

	fmt.Println("s append:", s)

	// Slicing is done by specifying a half-open range
	// with two indices separated by a colon
	fmt.Println("sl 2 to 4", s[2:4])
	fmt.Println("sl upto 3", s[:3])
	fmt.Println("sl from 2 to last", s[2:])
	fmt.Println()

	// Call the dummyFunc
	// Declare and initialize slices at the same time
	t := []int{1, 2}
	fmt.Println("slice as param:", dummyFunc(t))

	// Convert an array into a slice
	u := [3]string{"technicolor", "is", "saturated"}
	v := u[:]

	fmt.Println("array to slice:", v)
}
