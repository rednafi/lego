package main

import "fmt"
import "io"
import "net/http"

// Greet greets a name and prints it to the stdout
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler handles an http request to the port 5000
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
