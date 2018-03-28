package main

import "fmt"

func main() {
	fmt.Println("square(5):", square(5))
	fmt.Println("power(3, 4):", power(3, 4))
	fmt.Println("add(3, 4):", add(3, 4))
	fmt.Println("operation(3, 4, \"-\"):", operation(3, 4, "-"))
	plus, minus := plusAndMinus(6, 8)
	fmt.Println("plusAndMinus(6, 8):", plus, minus)
	fmt.Println("factorial(5):", factorial(5))
}

func square(a int) int {
	return a * a
}

func power(a int, b int) int {
	ret := 1
	for i := 0; i < b; i++ {
		ret *= a
	}
	return ret
}

// If it has same argument type, you can put it one at the end.
func add(a, b int) int {
	return a + b
}

// More complex.
func operation(a, b int, c string) int {
	switch c {
	case "+":
		return a + b
	case "-":
		return a - b
	default:
		return 0
	}
}

// Multiple return!
func plusAndMinus(a, b int) (int, int) {
	plus := a + b
	minus := a - b
	return plus, minus
}

// And finally, a recursion!
func factorial(a int) int {
	if a <= 0 {
		return 1
	}
	return a * factorial(a-1)
}
