/*
Channels can be thought as pipes using which Goroutines communicate. Similar to how water flows from one end to another in a pipe, data can be sent from one end and received from the another end using channels.

Each channel has a type associated with it. This type is the type of data that the channel is allowed to transport. No other type is allowed to be transported using the channel.

chan T is a channel of type T

The zero value of a channel is nil. nil channels are not of any use and hence the channel has to be defined using make similar to maps and slices.

Sending/ Receiving Data from Channels
--------------------------------------
data := <- a // read from channel a
a <- data // write to channel a

** Sends and receives are blocking by default **
Sends and receives to a channel are blocking by default. What does this mean? When a data is sent to a channel, the control is blocked in the send statement until some other Goroutine reads from that channel. Similarly when data is read from a channel, the read is blocked until some Goroutine writes data to that channel.
*/

package main

import "fmt"

// func hello(done chan bool){
// 	fmt.Println("Hello from goroutine!")
// 	done <- true
// }
// func main() {
// 	done := make(chan bool)
// 	go hello(done)
// 	<- done
// 	fmt.Println("Main function finished.")
// }

func calcSquare(num int, squarech chan int) {
	sum := 0

	for num != 0 {
		digit := num % 10
		sum += digit * digit
		num /= 10
	}
	// Instead of returning the value, you write to the channel
	// Later value written to this value will be retrieved
	squarech <- sum
}

func calcCube(num int, cubech chan int) {
	sum := 0

	for num != 0 {
		digit := num % 10
		sum += digit * digit * digit
		num /= 10
	}
	cubech <- sum
}

func main() {
	squarech, cubech := make(chan int), make(chan int)

	// These goroutines will write to their respective channels after execution
	go calcSquare(500, squarech)

	// We collect the values written by the goroutines
	squares, ok := <-squarech
	if ok != true {
		fmt.Println("Not received, data will be zero")
	}


	// Sending the sum of squares to calcCube
	go calcCube(squares, cubech)

	cubes, ok :=<- cubech
	if ok != true {
		fmt.Println("Not received, data will be zero")
	}

	fmt.Println("square sum:", squares)
	fmt.Println("cube sum:", cubes)

}
