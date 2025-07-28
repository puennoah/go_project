package main

import "fmt"

func main() {
	var miles = []int{10, 26, 45}
	var yards = []int{0, 385, 44}
	for i, _ := range miles {
		fmt.Println("Convert Miles and Yards to Kilometers :", miles[i], "mi.", yards[i], "yd.")
		fmt.Printf("Answer is %f kilometers.\n\n", convertToKm(miles[i], yards[i]))
	}
}

func convertToKm(miles int, yards int) float64 {
	return (1.609 * (float64(miles) + float64(yards)/1760.0))
}
