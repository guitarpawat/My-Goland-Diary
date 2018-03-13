package main

import "fmt"

func main() {
	// Basic switch-case
	// `default:` is equivalent to else.
	// And... you don't need to `break`!
	i := 2
	switch i {
	case 1:
		fmt.Println("You selected 1.")
	case 2:
		fmt.Println("You selected 2.")
	case 3:
		fmt.Println("You selected 3.")
	default:
		fmt.Println("Invalid number!")
	}

	// I am lazy!
	switch j := 3; j {
	// Multiple expression in case.
	case 1, 3, 5:
		fmt.Println("It's odd.")
	case 2:
		fmt.Println("You selected 2.")
	default:
		fmt.Println("Invalid number!")
	}

	// Switch without expression is alternative if.
	k := -2
	switch {
	case k > 0:
		fmt.Println("This is positive.")
	case k < 0:
		fmt.Println("This is negative.")
	case k == 0:
		fmt.Println("This is zero.")
	}

	// Switch with string
	s := "hello"
	switch s {
	case "hello":
		fmt.Println("Hello World")
	case "goodbye":
		fmt.Println("Goodbye World")
	}
}
