package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34

	for name, age := range ages {
		fmt.Printf("%s: %d\n", name, age)
	}
}
