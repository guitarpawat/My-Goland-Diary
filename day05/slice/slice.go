package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Create a slice with initial size 3.
	// make([]type,length)
	var a = make([]int, 3)
	fmt.Println(a)

	// Append value to slice.
	// Return new value, not append to old variable.
	var b = append(a, 1)
	b = append(b, 2, 3, 4, 5)
	fmt.Println(a)
	fmt.Println(b)

	// One line value declaration.
	var c = []int{1, 2, 3, 4}
	fmt.Println(c)

	// Copy the slice.
	var d = make([]int, 5)
	// copr(to,from)
	copy(d, c)
	fmt.Println(d)
	fmt.Println(len(d))

	// Access by index.
	d[4] = 5
	fmt.Println(d)
	fmt.Println(d[2])

	// Make 2D array.
	var e = make([][]int, 3)
	fmt.Println(e)
	for i := 0; i < 3; i++ {
		e[i] = make([]int, 3)
	}
	fmt.Println(e)

	// From array to slice.
	var f = [5]int{0, 1, 2, 3, 4}
	fmt.Println(reflect.TypeOf(f).Kind())
	var g = f[:]
	fmt.Println(reflect.TypeOf(g).Kind())

	// Slice an array.
	fmt.Println(f[:2])
	fmt.Println(f[2:])
	fmt.Println(f[1:3])
}
