package main

//import "fmt"

// Sum returns the summation of the integers in an array
func Sum(numbers []int) int {

	sum := 0
	for _, number := range numbers {
		sum += number

	}
	return sum
}

// SumAll ([]int{1,2}, []int{0,9}) would return []int{3, 9}
// or, SumAll([]int{1,1,1}) would return []int{3}
func SumAll(numbers ...[]int) []int {
	sums := []int{}
	for _, nums := range numbers {
		sum := Sum(nums)
		sums = append(sums, sum)

	}
	return sums
}

// func main() {
// 	fmt.Println(SumAll([]int{1, 2, 3}, []int{1}, []int{1}))

// 	// Declare and array
// 	s := make([]int, 5, 5)

// 	// Make a slice from the array
// 	t := s[:]

// 	// Append stuff to the array
// 	t = append(t, 1, 2, 3, 4, 5)
// 	fmt.Println(t)
// }
