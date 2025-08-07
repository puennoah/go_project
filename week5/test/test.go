package main

import "fmt"

func main() {

	x := 2

	y := -1

	z := 3

	fmt.Println("x is ", x, "   y is ", y, "  z is ", z)

        z++

	if x < y {

		fmt.Println("x is ", x, "   y is ", y, "  z is ", z)

		x = y

	} else {

		fmt.Println("x is ", x, "   y is ", y, "  z is ", z)

            z  += 2

		y++

	}

	if z := 9; z < x {

		fmt.Println("x is ", x, "   y is ", y, "  z is ", z)

		y--

	} else {

		fmt.Println("x is ", x, "   y is ", y, "  z is ", z)

                z *= 2

	}

	fmt.Println("x is ", x, "   y is ", y, "  z is ", z)

}