package main

// problem statement
type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

type GoCircle struct {
	Point
	Radius int
}

type GoWheel struct {
	GoCircle
	Spokes int
}

func main() {
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20

	var gow GoWheel
	gow.X = 8
	gow.Y = 8
	gow.Radius = 5
	gow.Spokes = 20
}
