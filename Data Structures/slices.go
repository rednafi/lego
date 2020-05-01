package main

import "fmt"

/*
- Slices are similar to arrays, only more powerful
- Their size is not fixed and can be updated dynamically
- Indexing starts from 0
- However, while dynamically increasing or decreasing their size, slices use more memrory. So if you know the size of your container, arrays are a better choice. Use slices judiciously.
*/

// Declaring slices in a function
func dummyFunc(t []int) []int {
	t = append(t, 5)
	t = append(t, 4)
	return t
}

func main() {
	// Declaring a slice using the built-in make function

	s := make([]string, 3)

	fmt.Println("s empty:", s)

	s[0] = "Go"
	s[1] = "is"
	s[2] = "awesome"

	fmt.Println("s init:", s)

	// Append new values to the end of slice
	s = append(s, "let's")
	s = append(s, "go")

	fmt.Println("s append:", s)

	// Slice and dice
	fmt.Println("sl 2 to 4", s[2:4])
	fmt.Println("sl upto 3", s[:3])
	fmt.Println("sl from 2 to last", s[2:])
	fmt.Println()

	// Call the dummyFunc
	// Declaring and initializing slices at the same time
	t := []int{1, 2}
	fmt.Println("slice as param:", dummyFunc(t))
}
