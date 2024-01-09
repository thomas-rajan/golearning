package main

import "fmt"

func swap() {
	a := 100
	b := 200
	fmt.Printf("a = %v, b = %v\n", a, b)
	temp := a
	a = b
	b = temp
	fmt.Printf("a = %v, b = %v\n", a, b)
}

func swapWithTuple() {
	a := 100
	b := 200
	fmt.Printf("swapWithTuple a = %v, b = %v\n", a, b)
	b, a = a, b
	fmt.Printf("swapWithTuple a = %v, b = %v\n", a, b)
}

func main() {
	swap()
	swapWithTuple()
}
