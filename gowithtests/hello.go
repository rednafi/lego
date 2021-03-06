package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello func prints to stdout
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix

	case "French":
		prefix = spanishHelloPrefix
	}

	return prefix + name

}

func main() {
	fmt.Println(Hello("", "Spanish"))
}
