package main

import (
	"fmt"
	"math"
)

func main() {
	var area, radius float32
	fmt.Println("Input radius in meters")
	fmt.Scanf("%f", &radius)
	area = math.Pi * radius * radius
	fmt.Printf("radius of %g meters; area is %g square meters\n", radius, area)
}
