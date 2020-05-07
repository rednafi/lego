// package main

import (
	"errors"
	"fmt"
)

// Return a simple error with built in errors interface
func someFunc(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("A single integer can't be the answer for everything")
	}
	return arg + 3, nil
}

// Write a custom error
// For this you will need to implement Error method on a custom argError struct
type argError struct {
	arg  int
	prob string
}

// Define the Error method
// While fmt.Prinf simply prints a formatted string,
// fmt.Sprint creates the string and then returns it
func (e *argError) Error() string {
	return fmt.Sprintf("arg: %d, problem: %s", e.arg, e.prob)
}

// Use the method now to raise a custom error
func anotherFunc(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg: arg, prob: "A single integer can't be the answer for everything."}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {

		// Result from someFunc
		fmt.Println("\nResult from someFunc:")
		res, err := someFunc(i)
		if err != nil {
			fmt.Println(res, err)
		} else {
			fmt.Println(res)
		}

		// Result from anotherFunc
		fmt.Println("\nResult from anotherFunc:")
		res1, err1 := anotherFunc(i)
		if err1 != nil {
			fmt.Println(res1, err1)
		} else {
			fmt.Println(res1)
		}
	}

	// Another idiom to handle error (anotherFunc)
	// If you want to programmatically use the data in a custom error,
	// you’ll need to get the error as an instance of the custom error
	// type via type assertion.
	fmt.Println("\nResult from last handle:")
	_, err := anotherFunc(42)
	ae, _ := err.(*argError)
	fmt.Println(ae.arg)
	fmt.Println(ae.prob)
}
