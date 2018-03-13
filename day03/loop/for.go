package main

import "fmt"

func main() {
	i := 5
	// while(i > 0)
	for i > 0 {
		fmt.Println(i)
		i--
	}

	// while(true)
	for {
		fmt.Println("It is dangerous without break!")
		// Stop it or an infinite loop!
		break
	}

	// for(int j=0; j<5; j++)
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	// Continue for the next iteration
	for k := 0; k < 5; k++ {
		fmt.Println("Hello", k)
		if k == 2 {
			continue
		}
		fmt.Println("Hi", k)
	}

	// Cannot use var to declare within for
	// for var m = 1; m <5; m++ {

	// }
}
