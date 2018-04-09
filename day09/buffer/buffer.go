package main

import (
	"fmt"
)

func main() {

	// Create a channel with buffer size 2.
	msg := make(chan string, 2)

	// Send data to the channel using `<-`.
	go func() { msg <- "hello" }()

	go func() { msg <- "world" }()

	// Get the data using `<-`.
	// Spot the difference!
	str := <-msg
	fmt.Println(str)

	str = <-msg
	fmt.Println(str)

	// Error.
	// str = <-msg
	// fmt.Println(str)
}
