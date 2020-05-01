package main

import "fmt"

/*
- Arrays have fixed length
- They are zero indexed
- Length of an array can be derived with len function.
- Arrays types are 1D but high dimensional arrays can be formed using 1D arrays
*/

func main() {
	// Declaring an array of integers with length 5
	var a [5]int

	// Filling up array by index
	a[0] = 1
	a[2] = 2

	// Declaring and initializing with composite literal
	b := [4]string{"hello", "from", "the", "compiler"}

	// Filling up array with a for loop
	var c [3]float64

	for i := 0; i < 3; i++ {
		// Converting int to float
		j := float64(i)
		c[i] = j
	}

	// Forming a 2D array (2 x 3 array)
	var arr2d [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			arr2d[i][j] = i + j
		}
	}
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("arr2d:", arr2d)

}
