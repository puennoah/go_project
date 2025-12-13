package main

import (
	"fmt"
)

type Number interface {
	string | complex64 | int | int8 | int16 | int32 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64
}

func main() {
	// Example usage with int slice
	intSlice := []int{1, 2, 3, 4, 5}
	intSum := Sum(intSlice)
	fmt.Printf("Sum of int slice: %d\n", intSum)

	// Example usage with float64 slice
	floatSlice := []float64{1.2, 3.4, 5.6, 7.8}
	floatSum := Sum(floatSlice)
	fmt.Printf("Sum of float64 slice %.2f\n", floatSum)

	// Example usage with uint16 slice
	uintSlice := []uint16{10, 20, 30, 40}
	uintSum := Sum(uintSlice)
	fmt.Printf("Sum of uint16 slice: %d\n", uintSum)

	// Example usage with string slice
	strSlice := []string{"if", " pigs", " can", " fly."}
	strSum := Sum(strSlice)
	fmt.Printf("Sum of string slice: %s\n", strSum)

	// Example usage with complex64 slice
	complexSlice := []complex64{1.2 +8.8i, 3.4, 5.6, 7.8}
	complexSum := Sum(complexSlice)
	fmt.Printf("Sum of complex slice: %.4v\n", complexSum)

}

func Sum[T Number](slice []T) T {
	var sum T
	for _, v := range slice{
		sum += v
	}
	return sum
}