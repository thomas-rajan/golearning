package main

import "fmt"

func main() {
	var values [3]int
	fmt.Println("element at index 1, ", values[1])

	for index, value := range values {
		fmt.Printf("element at index %v is %v\n", index, value)
	}
	// array initialization
	var q [3]int = [3]int{1, 2, 3}

	r := [...]int{1, 2, 3}

	fmt.Println("q element at index 1, ", q[1])

	fmt.Println("r element at index 1, ", r[1])
}
