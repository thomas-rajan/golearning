package main

type Employee struct {
	ID     int
	Name   string
	Salary int
}

func main() {
	e := Employee{
		ID:     1,
		Name:   "Alice",
		Salary: 100,
	}
	e.Salary = 200
}
