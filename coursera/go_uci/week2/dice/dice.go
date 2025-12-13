package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var pair, vPair, howMany int
	fmt.Println("reading pair of dice value")
	fmt.Scanf("%d", &vPair)
	for i := 0; i < 10000000; i++ {
		pair = rand.Intn(6) + rand.Intn(6) + 2
		if pair == vPair {
			howMany++
		}
	}
	fmt.Println(howMany)
	fmt.Printf("Probability of rolling a %d is %4g \n", vPair, float32(howMany)/10000000.0)
}
