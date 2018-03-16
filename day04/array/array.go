package main

import "fmt"

func main() {
	// Initialize an array with length 3.
	// You must tell the type of array.
	// var a [3]
	var a [3]int
	fmt.Println(a)
	// Get the length of array.
	fmt.Println(len(a))
	// Assign a value into array.
	a[1] = 2
	fmt.Println(a)

	// One line initialize array value
	b := [3]int{2, 3, 7}
	fmt.Println(b)

	// 2D array.
	c := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	// or `var c [2][3]int`
	fmt.Println(c)
}
