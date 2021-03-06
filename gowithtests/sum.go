package main

// Sum returns the summation of the integers in a slice
func Sum(numbers []int ) int {

	sum := 0
	for _, number := range numbers{
		sum += number

	}
	return sum
}
