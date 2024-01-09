package main

import "fmt"

func main() {
	x := 100
	p := &x

	fmt.Printf("Before x = %v, *p = %v \n", x, *p)

	*p = 200
	fmt.Printf("After x = %v, *p = %v \n", x, *p)
}
