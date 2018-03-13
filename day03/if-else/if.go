package main

import "fmt"

func main() {
	// Basic if statement.
	i := 5
	if i%2 == 0 {
		fmt.Println("5 is divisible by 2.")
	} else {
		fmt.Println("5 is not divisible by 2.")
	}

	// else required to be the same line with `}`.
	// if i%2 == 0 {
	// 	fmt.Println("5 is divisible by 2.")
	// }
	// else {
	// 	fmt.Println("5 is not divisible by 2.")
	// }

	// Cannot
	// if i%2 == 0 fmt.Println("5 is divisible by 2.")
	// else fmt.Println("5 is not divisible by 2.")

	// Multiple if-else
	if i < 0 {
		fmt.Println("i is less than 0.")
	} else if i > 0 {
		fmt.Println("i is more than 0.")
	} else {
		fmt.Println("i is equals to 0")
	}

	// Declare variable within in if-else statement
	if n := 1; n < 0 {
		fmt.Println("n is less than 0.")
	} else if n > 0 {
		// Also valid
		fmt.Println("n is more than 0.")
		fmt.Println(n)
	} else {
		fmt.Println("n is equals to 0")
	}

	// Invalid.
	// fmt.Println(n)

	// == operator is valid for compare strings.
	s := "hello"
	if s == "hi" {
		fmt.Println("Why???")
	} else if s == "hello" {
		fmt.Println("Yay!!!")
	}
}
