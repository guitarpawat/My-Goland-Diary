package main

import "fmt"

func main() {
	// Variable declaration.
	var a = 1
	fmt.Println(a)
	// Cannot change type
	// a = "hello"

	// Assign multiple variable at once.
	var b, c = 1 + 1, "hello"
	fmt.Println(b)
	fmt.Println(c)

	// Declare with type.
	var d string = "hello"
	fmt.Println(d)

	// Default value for each type.
	var e string
	var f int
	var g float64
	fmt.Println(e) //""
	fmt.Println(f) //0
	fmt.Println(g) //0

	// Variable declaration shorthand.
	h := "hi"
	fmt.Println(h)
	// But don't do this for reassign.
	// h := "hello"
}
