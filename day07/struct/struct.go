package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	s1 := person{"Guitar", 19}
	fmt.Println(s1)

	// Specify variable while initialize.
	s2 := person{name: "Pawat", age: 23}
	fmt.Println(s2)

	// Omitted value will be default value.
	s3 := person{name: "gp"}
	fmt.Println(s3)

	// Copy struct.
	s4 := s1
	fmt.Println(s4)

	// Copy pointer
	s5 := &s1
	fmt.Println(s5)

	// Get value of struct.
	fmt.Println(s1.name)

	// Set value of struct.
	s1.name = "hello"
	fmt.Println(s1.name)
	fmt.Println(s4.name)
	// Value changed!
	fmt.Println(s5.name)
}
