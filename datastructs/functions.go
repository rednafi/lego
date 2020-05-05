package main

import (
	"fmt"
	"math"
)

// Define a function
func someAdd(a, b int) int {
	return a + b
}

// Define a variadic function
// Variadic functions take any number of arguments
func variadicAdd(a ...int) int {
	sum := 0
	for _, num := range a {
		sum += num
	}
	return sum
}

// Return multiple values
func someDiv(a, b int) (float64, bool) {

	if b == 0 {
		res, ok := math.Inf(0), false
		return res, ok
	}
	res, ok := float64(a/b), true
	return res, ok

}

func main() {
	fmt.Println("result someAdd:", someAdd(1, 2))
	fmt.Println("result variadicAdd:", variadicAdd(0, 100, 200))
	res, ok := someDiv(1, 0)
	fmt.Println("result someDiv:", res, ok)
}
