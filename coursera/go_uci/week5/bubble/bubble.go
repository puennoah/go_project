package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("data of Int random numbers 0 to 10000")

	size := 3000
	dataSlice := make([]int, size)
	for i, _ := range dataSlice {
		dataSlice[i] = rand.Intn(10000)
	}
	bubbleSort(dataSlice)
	fmt.Println(dataSlice[:20])
	fmt.Println(dataSlice[size-20:])
}

func swap (a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}


func bubbleSort(d []int) {
	for i := 0; i < len(d) - 1; i++ {
		for j := 0; j <  len(d) - i - 1; j++ {
			if d[j] > d[j+1] {
				swap(&d[j], &d[j+1])
			}
		}
	}
}