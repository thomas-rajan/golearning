package main

import "fmt"

const globalX = 100

func printGlobal() {
	fmt.Printf("from printGlobal globalX = %v \n", globalX)
}

func main() {
	localX := 200
	localX = 300
	fmt.Printf("from main globalX = %v, localX = %v \n", globalX, localX)

	printGlobal()
}
