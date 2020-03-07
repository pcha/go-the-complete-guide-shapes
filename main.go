package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base float64
}

func main() {
	s := square{sideLength:5}
	t := triangle{base:5, height:6}

	fmt.Println("The area of the square of side ", s.sideLength, " is ", s.getArea())
	fmt.Println("The area of the triangle of base ", t.base, " and height ", t.height, " is ", t.getArea())
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}

func (t triangle) getArea() float64 {
	return t.base * t.height / 2
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}