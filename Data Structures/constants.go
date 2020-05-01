package main

import "fmt"

// Constants can be defined just like variables but they are immutable

// Multiple constants
const a, b int = 1, 2

// Another way
const (
	pi float64 = 3.141
	phi float64 = 1.618
	e float64 = 2.718
)

func main() {
	fmt.Println("a:", a, "b:", b)

	// Print the types
	fmt.Printf("Type of a is %T and type of b is %T", a, b)
}

