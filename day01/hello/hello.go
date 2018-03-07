package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	//Cannot do this.
	//fmt.Println("Hello" + 1)
	fmt.Println("Hello", 1+1)
	fmt.Println(1+1, 2)

	//It is valid if they are both string.
	fmt.Println("Hello " + "World")
	//Also valid.
	fmt.Println("Hello", "World")

	//Just a basic boolean.
	fmt.Println(true || false)
	fmt.Println(true && false)
	fmt.Println(!false)
}
