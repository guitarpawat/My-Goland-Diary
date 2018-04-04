package main

import (
	"errors"
	"fmt"
)

func summation(i int) (int, error) {
	if i < 0 {
		// Error, golang has no exception!
		return -1, errors.New("Number must be zero or positive")
	}

	ret := 0

	for n := 0; n <= i; n++ {
		ret += n
	}

	// No error.
	return ret, nil
}

func printResult(ret int, err error) {
	if err == nil {
		fmt.Println("Result:", ret)
	} else {
		fmt.Println("Error:", err)
	}
}

func main() {

	printResult(summation(-10))

	printResult(summation(10))
	printResult(summation(-0))
}
