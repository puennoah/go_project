package main

import (
	"fmt"
	"math/rand"
)

const doy int = 365 //days of year
const sim int = 10000

func main() {
	fmt.Println("People in the room\t\tProbability of 2 having the same birthday")
	for people := 10; people <= 100; people += 10 {
		countSameBirthday := 0
		for i := 0; i < sim; i++ {
			seen := make([]bool, 365)
			for j := 0; j < people; j++ {
				day := rand.Intn(doy)
				if seen[day] == true {
					countSameBirthday++
					break
				} else {
					seen[day] = true
				}

			}

		}
		fmt.Printf("%d\t\t\t\t%.3f\n",people, float64(countSameBirthday)/float64(sim))
	}
}
