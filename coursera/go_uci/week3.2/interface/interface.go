package main

import (
	"fmt"
)




type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.12 * c.radius * c.radius
}

type Square struct {
	side float64
}

func (sq Square) area() float64 {
	return sq.side * sq.side
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func main() {
	c := Circle{radius: 5}
	sq := Square{side: 6}
	r := Rectangle{width: 3, height: 4}

	shapes := []Shape{c, sq, r}
	for _, shape := range shapes {
		fmt.Printf("Area of the shape:  %f\n", shape.area())
	}
}




