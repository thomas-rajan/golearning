package main

import "fmt"

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(n, m int) int { return n * m }

func main() {
	f := square
	fmt.Println(f(3)) //9

	f = negative
	fmt.Println(f(3)) //-3

	//g := product //compiler error
}
