package main

import "fmt"

func main() {
	// range iterate over arrays, slices and maps.
	s1 := []string{"hi", "hello", "nihao"}
	for index, value := range s1 {
		fmt.Println(index, ":", value)
	}

	s2 := map[string]string{"jpn": "japan", "eng": "english", "usa": "america"}
	for key, value := range s2 {
		fmt.Println(key, ":", value)
	}

	// Iterate only key.
	for key := range s2 {
		fmt.Println(key)
	}

	// Iterate only value.
	for _, value := range s2 {
		fmt.Println(value)
	}

	// range iterate string as rune (unicodes).
	for index, value := range "golang" {
		fmt.Println(index, ":", value)
	}
}
