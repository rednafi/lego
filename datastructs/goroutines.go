package main

import (
	"fmt"
	"strings"
	"time"
)

/*
Goroutines are comparable to threads, only a single kernal thread can house hundreds or even thousands of goroutines. Goroutines have lower memory footprint and communication overhead than os threads.

To start a goroutine, you just define a normal function and add 'go' prefix before calling the function.

If you omit the time.Sleep statement in the main function, the first function
will not print anything. This happens because

	* When a new Goroutine is started, the goroutine call returns immediately. Unlike functions, the control does not wait for the Goroutine to finish executing. The control returns immediately to the next line of code after the Goroutine call and any return values from the Goroutine are ignored.

	* The main Goroutine should be running for any other Goroutines to run. If the main Goroutine terminates then the program will be terminated and no other Goroutine will run.

This way of using sleep in the main Goroutine to wait for other Goroutines to finish their execution is a hack we are using to understand how Goroutines work. Channels can be used to block the main Goroutine until all other Goroutines finish their execution
*/

func hello() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Hello from goroutine!")

}

// Running multiple goroutines
func numbers() {
	for i := 0; i < 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func alphabets() {
	for _, i := range strings.Split("abcde", "") {
		time.Sleep(450 * time.Millisecond)
		fmt.Printf("%s", i)
	}
}

func main() {
	go hello()
	go numbers()
	go alphabets()

	// THis is a hack to stop running the main goroutine
	// before hello goroutine finishes executing.
	time.Sleep(2000 * time.Second)
	fmt.Println("Main function terminated")
}
