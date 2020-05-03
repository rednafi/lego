package main

import "fmt"

// This is valid both inside and outside of functions
var a int = 1

// Defining multiple variables
var b, c int = 2, 3

// Another way to reduce boiler plate
var (
	integerVar int     = 4
	stringVar  string  = "Go"
	floatVar   float64 = 5.7
)

// Defining & initializing variable with := operator
// This is only valid inside functions
func dummyFunc() (int, string, float64) {
	dummyInt := 5 //Compiler infers the type
	dummyString := "tekashi69"
	dummyFloat := 56.9
	return dummyInt, dummyString, dummyFloat
}

func main() {
	fmt.Println("a:", a)
	fmt.Println("b:", b, "c:", c)
	fmt.Println("integerVar:", integerVar, "stringVar:",
		stringVar, "floatVar:", floatVar)

	dummyInt, dummyString, dummyFloat := dummyFunc()
	fmt.Println("dummyInt:", dummyInt, "dummyString:", dummyString, "dummyFloat:", dummyFloat)
}
