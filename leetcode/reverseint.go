/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21

*/
package main

import (
	"fmt"
	"strconv"
	"strings"
	"math"
)

func reverse(x int) int {

	// Check sign
	sign := 1
	if x < 0 {
		sign = -1
	}

	// Take absolute value
	if sign == -1 {
		x = -1 * x
	}

	// Convert int to str
	xStr := strconv.Itoa(x)

	// Convert str to slice
	xSlice := strings.Split(xStr, "")

	// Reverse slice
	i := 0
	j := len(xSlice) - 1

	for i < j {
		xSlice[i], xSlice[j] = xSlice[j], xSlice[i]
		i++
		j--
	}

	// Slice to string
	xStr = strings.Join(xSlice, "")

	// String to int
	x, _ = strconv.Atoi(xStr)
	x = sign * x

	// Check if reverse integer overflows
	if x < - int(math.Pow(2,31)) || x > int(math.Pow(2,31))-1 {
		return 0
	}

	return x

}

func main() {
	fmt.Println(reverse(1563847412))
}
