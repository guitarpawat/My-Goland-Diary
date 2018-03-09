package main

import "fmt"
import "reflect"

func main() {
	var a = 3
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))

	var b = int(a)
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))

	var c = 5 / 2
	fmt.Println(c)
	fmt.Println(reflect.TypeOf(c))

	// Does not matter it cast
	var d = float64(c)
	fmt.Println(d)
	fmt.Println(reflect.TypeOf(d))

	var e = 3.99
	fmt.Println(e)
	fmt.Println(reflect.TypeOf(e))

	// Casting round it down.
	var f = int(e)
	fmt.Println(f)
	fmt.Println(reflect.TypeOf(f))

	// Cannot do this.
	// var g = "H"
	// var h = int(g)

	var i = 't'
	fmt.Println(i)
	fmt.Println(reflect.TypeOf(i))
	// Use Printf instead!
	fmt.Printf("%c\n", i)

	var j = byte(i)
	fmt.Println(j)
	fmt.Println(reflect.TypeOf(j))

	var k = true
	fmt.Println(k)
	fmt.Println(reflect.TypeOf(k))

	// Cannot convert bool to int.
	// var l = int(k)
	// fmt.Println(l)

	// And string.
	//var m = string(k)

	// Also cannot do.
	// var n = 1
	// var o = string(n)
}
