package main 

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("size of slice is ")
	size := 0
	fmt.Scanf("%d", &size)
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Float64() * 1000.0
		fmt.Println(data[i])
	}
	fmt.Println("min of", size, "rand.Float64 is", minSlice(data))
}

func minSlice(d []float64) float64 {
	var min float64 = 1000.0
	for _, v := range d {
		if v < min {
			min = v
		}
	}
	return min
}