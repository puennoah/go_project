package main

import (
	"errors"
	"fmt"
	"math"
)

func Roots(a, b, c float64) (float64, float64, error) {
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return 0, 0, errors.New("discriminant is negative")
	}
	d := math.Sqrt(discriminant)
	return (-b +d) / (2 * a), (-b -d) / (2 * a), nil
}

func main() {
	root1, root2, err := Roots(1.0, 2.0, 1.0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: root1 =  %f root2 = %f\n", root1, root2)
	}

	root3, root4, err := Roots(1.0, 2.0, 4.0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: root3 =  %f root4 = %f\n", root3, root4)
	}
}