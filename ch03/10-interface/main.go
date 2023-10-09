package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func PrintArea(s Shape) {
	fmt.Printf("%T's area is %0.2f\n", s, s.Area())
}

func main() {
	r := Rectangle{Width: 5, Height: 10}
	PrintArea(r)

	c := Circle{Radius: 5}
	PrintArea(c)

	var s Shape
	s = Circle{}
	C := s.(Circle)
	fmt.Printf("%T\n", C)
}
