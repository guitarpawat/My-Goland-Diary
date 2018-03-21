package main

import "fmt"

func main() {
	// Create the map variable.
	// make(map[key type]value type)
	a := make(map[string]int)
	fmt.Println(a)
	fmt.Println(len(a))

	// Add data to the map.
	a["pencil"] = 5
	a["book"] = 30
	fmt.Println(a)
	fmt.Println(len(a))

	// No book for today!
	delete(a, "book")
	fmt.Println(a)
	fmt.Println(len(a))

	// Getting the value
	// reurns value, has key
	b, c := a["pencil"]
	fmt.Println(b)
	fmt.Println(c)

	d, e := a["book"]
	fmt.Println(d) // Default value of int is 0.
	fmt.Println(e) // Does not have this value.

	// Usage
	if e {
		fmt.Println(d)
	}

	// One line declaration.
	f := map[string]int{"one": 1, "two": 2}
	fmt.Println(f)
}
