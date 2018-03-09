package main

import "fmt"
import "reflect"

func main() {
	// To declares the constant.
	const a = 1
	fmt.Println(a)
	const b = 3.14
	fmt.Println(b)
	const c = "Hello"
	fmt.Println(c)

	// Cannot do like this.
	// const d
	// fmt.Println(d)

	// Must initialize the value when declare.
	// const e
	// e = 5
	// fmt.Println(e)

	const f = 5.0 / 2
	fmt.Println(f)
	fmt.Println(reflect.TypeOf(f))

	const g = 5 / 2
	fmt.Println(g)
	fmt.Println(reflect.TypeOf(g))

}
