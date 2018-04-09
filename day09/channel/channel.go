package main

import (
	"fmt"
)

func main() {

	// Create a channel.
	msg := make(chan string)

	// Send data to the channel using `<-`.
	go func() { msg <- "hello" }()

	go func() { msg <- "world" }()

	// Get the data using `<-`.
	// Get lastest first.
	str := <-msg
	fmt.Println(str)

	str = <-msg
	fmt.Println(str)

	// Error.
	// str = <-msg
	// fmt.Println(str)
}
