package main

import (
	"fmt"
)

// type bot interface {
// 	getGreeting() string
// }

// type englishBot struct{}
// type spanishBot struct{}

// func main() {
// 	eb := englishBot{}
// 	sb := spanishBot{}

// 	printGreeting(eb)
// 	printGreeting(sb)
// }

// func printGreeting(b bot) {
// 	fmt.Println(b.getGreeting())
// }

// func (englishBot) getGreeting() string {
// 	// VERY custom logic for generating an english greeting
// 	return "Hi There!"
// }

// func (spanishBot) getGreeting() string{
// 	return "Hola!"
// }

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	tri := triangle{
		height: 5.0,
		base:   5.0,
	}

	squ := square{
		sideLength: 5.0,
	}
	printArea(squ)
	printArea(tri)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println("Area is ", s.getArea())
}
