package main

import "fmt"

type class struct {
	name     string
	students int
}

func makeClass(name string) *class {
	c := class{name: name}
	c.students = 100
	return &c
}

func main() {
	// Construct the struct
	fmt.Println("raw struct:", class{name: "firstG", students: 100})

	// Construct the struct via function
	fmt.Println("from function", makeClass("secondG"))

	// Half initialize
	k := class{name: "thirdG"}
	fmt.Println("half initialize:", k )
	fmt.Println("empty attribute:", k.students)

	//You can also use dots with struct pointers -
	//the pointers are automatically dereferenced.
	l := &k
	fmt.Println("as pointer:", l.name )

	//Mutate
	l.name = "fourthG"
	fmt.Println("mutate:", l)
}
