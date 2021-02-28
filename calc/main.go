package main

import (
	"calc/addsub"
	"calc/muldiv"
	"fmt"
)

// CheckError does rudimentary error checking.
func CheckError(e error) interface{} {
	if e != nil {
		return fmt.Errorf("an error happened")
	}
	return nil
}

func main() {

	r, err := addsub.Add(20, 44)
	CheckError(err)
	fmt.Printf("Value of add func is %v \n", r)

	r, err = addsub.Sub(20, 44)
	CheckError(err)
	fmt.Printf("Value of sub func is %v \n", r)

	r, err = muldiv.Mul(20, 44)
	CheckError(err)
	fmt.Printf("Value of mul func is %v \n", r)

	s, err := muldiv.Div(200, 44)
	CheckError(err)
	fmt.Printf("Value of div func is %v \n", s)
}
