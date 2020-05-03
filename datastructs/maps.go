package main

import "fmt"

func main() {

	// Make an empty map where the keys are strings
	// & the values are slices of integers
	m := make(map[string][]int)
	fmt.Println("empty map:", m)

	// Assign values to the map
	m["list1"] = []int{0, 1, 2, 3}
	m["list2"] = []int{4, 5, 6, 7}

	// Print the map after assignment
	fmt.Println("original map", m)

	// Print length of the map (number of keys)
	fmt.Println("length of map:", len(m))

	// Delete
	delete(m, "list1")
	fmt.Println("deleted list1:", m)

	// Check if key list1 exists
	_, exists := m["list1"]
	fmt.Println("list1 exists:", exists)

	//Initialize map with composite literal
	n := map[string]int{"num1": 0, "num2": 1, "num3": 2}
	fmt.Println(n)
}
