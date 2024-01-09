package main

import "fmt"

func f() {}

func f1() int { return 100 }

var g = "g"

func scopeForLoop() {
	x := "hello"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}
}

func scopeIfLoop() {
	if x := f1(); x > 80 {
		fmt.Println("distinction", x)
	} else if x > 60 {
		fmt.Println("first class", x)
	} else {
		fmt.Println("failed", x)
	}
	fmt.Println(x) //compiler error
}

func main() {
	f := "f"
	fmt.Println(f) //local variable f shadows package level func f
	fmt.Println(g) //package level var
	//fmt.Println(h) //compiler error: undefined: h
}
