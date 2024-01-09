package main

import "fmt"

type Celsius float64
type Farenheit float64

func setFurnaceTemperature(f Farenheit) {
	fmt.Println("furnace temperature set to ", f)
}

func main() {
	far := Farenheit(100.0)
	setFurnaceTemperature(far)

	cel := Celsius(100.0)
	cel1 := Farenheit(cel)
	setFurnaceTemperature(cel1)
}
