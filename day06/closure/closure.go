package main

import "fmt"

func main() {
	// Create a closure of function.
	f1 := square()
	fmt.Println("f1(3):", f1(3))
	fmt.Println("f1(5):", f1(5))
	// 3+b
	plus3 := plus(3)
	fmt.Println("plus3(5):", plus3(5))
	fmt.Println("plus3(9):", plus3(9))
	// 9+b
	plus9 := plus(9)
	fmt.Println("plus9(5):", plus9(5))
	fmt.Println("plus9(9):", plus9(9))

	// Anonymous function
	hello := func(name string) string {
		return "Hello " + name
	}
	fmt.Println("hello(\"world\"):", hello("world"))
	fmt.Println("hello(\"moon\"):", hello("moon"))
}

// Function that returns function!
func square() func(int) int {
	return func(a int) int {
		return a * a
	}
}

// More complex!
func plus(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}
