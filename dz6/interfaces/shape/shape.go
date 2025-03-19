package shape

import (
	"fmt"
)

// Shape интерфейс для фигур
type Shape interface {
	// Area возвращает площадь фигуры
	Area() float64
}

// Circle реализует интерфейс Shape
type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Rectangle реализует интерфейс Shape
type Rectangle struct {
	Width, Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

// PrintArea выводит площадь фигуры
func PrintArea(s Shape) {
	fmt.Println(s.Area())
}
