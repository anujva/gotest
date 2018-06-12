package main

import (
	"fmt"
	"math"
)

type triangle struct {
	sideA float64
	sideB float64
	sideC float64
}

type square struct {
	side float64
}

type shape interface {
	getArea() float64
}

func main() {
	s := square{5}
	t := triangle{4.2, 3, 5}
	fmt.Println("Area of square:")
	printArea(s)
	fmt.Println()
	fmt.Println("Area of triangle")
	printArea(t)
}

func (t triangle) getArea() float64 {
	halfSum := (t.sideA + t.sideB + t.sideC) / 2
	return math.Sqrt(halfSum * (halfSum - t.sideA) * (halfSum - t.sideB) * (halfSum - t.sideC))
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
