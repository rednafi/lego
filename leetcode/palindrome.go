/*
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:

Input: 121
Output: true
Example 2:

Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*/

package main

import ("fmt"
		"strings"
		"strconv")

func isPalindrome(x int) bool {

	// If negative then false
	if x<0{
		return false
	}

	// Reverse number
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
	y, _ := strconv.Atoi(xStr)

	if x == y{
		return true
	}
	return false
}

func main(){
	fmt.Println(isPalindrome(-12))
}
