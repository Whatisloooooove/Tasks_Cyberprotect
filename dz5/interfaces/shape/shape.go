package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
	Weight, Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Weight
}

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

func main() {
	circle := Circle{Radius: 25}
	rectangle := Rectangle{Weight: 10, Height: 5}

	PrintArea(&circle)
	PrintArea(&rectangle)
}
