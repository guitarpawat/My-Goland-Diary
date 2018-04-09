package main

import (
	"fmt"
)

func count(s string, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(s+":", i)
	}
}

func hello() {
	fmt.Println("Hello")
}

func main() {
	// Running in normal way.
	count("count0", 10)

	fmt.Println("count0 ended")

	// Running concurrent asynchronously.
	go count("count1", 100)
	go count("count2", 100)

	fmt.Scanln()

	go count("count3", 100)
	go hello()

	fmt.Scanln()
}
