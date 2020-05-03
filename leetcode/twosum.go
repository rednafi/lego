/*Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

package main

import "fmt"

func twoSum(nums []int, target int) []int {

	for i, num1 := range nums {
		for j, num2 := range nums {
			if i != j {
				k := num1 + num2
				if k == target {
					indices := []int{i, j}
					return indices
				}
			}
		}
	}
	var indices []int
	return indices
}

func main() {
	nums := []int{ 3, 3}
	indices := twoSum(nums, 6)
	fmt.Println(indices)
}
