package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Blank identifier is use when the function
	// returns value but you don't need it.
	// It can be assign a value but never read from it
	// and complier also ignore it too!

	// We need to receive input from stdin without
	// declare the error variable
	reader := bufio.NewReader(os.Stdin)
	// text, err := reader.ReadString('\n')
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", 1)
	fmt.Println(text)

	// We need to check that the map has key or not.
	mapval := map[string]int{"one": 1, "two": 2}
	_, has := mapval["ten"]
	fmt.Println(has)

	// For debugging.
	// The unused import is considered ae an error.
	// Use blank identifier to ignore.
	_ = log.Flags // For debugging; delete when done.
}
