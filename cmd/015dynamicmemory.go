package main

import "fmt"

func main() {
	a := new(int)
	*a = 100
	fmt.Println(*a)
}
