package main

import "fmt"

func main() {
	i := 0
	fmt.Println("Initial value:", i)
	setVal(i)
	fmt.Println("setVal(i):", i)
	setPtr(&i)
	fmt.Println("setPtr(&i):", i)
	fmt.Println("Address of i", &i)
	// Pass the copy of i to function.
	setVal(i)
	fmt.Println("setVal(i):", i)
	// Pass the pointer of i to function.
	setPtr(&i)
	fmt.Println("setPtr(&i):", i)
	fmt.Println("Address of i", &i)
	j := i
	fmt.Println("Value of j:", j)
	fmt.Println("Address of j", &j)
	// Declare k to address of i.
	k := &i
	fmt.Println("Value of k:", k)
	fmt.Println("Address of k", &k)
}

func setVal(val int) {
	val++
}

func setPtr(ptr *int) {
	*ptr++
}
