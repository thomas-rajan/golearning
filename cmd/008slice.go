package main

import "fmt"

func main() {
	months := [...]string{"", "jan", "feb", "march", "april", "may", "june", "july", "aug", "sept", "octo", "nov", "dec"}

	q2 := months[4:7]
	fmt.Println("q2[0] = ", q2[0])
	/*
		summer := months[6:9]

		allMonths := months[1:]

		all := months[:]
	*/
}
