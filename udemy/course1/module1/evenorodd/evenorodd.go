package main

import (
	"fmt"
)

func main() {
	intSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10} 
	// Could improve and create with loop


	for _, v := range intSlice {
		if v % 2 == 0 {
			fmt.Println(v,"is even")
		} else {
			fmt.Println(v, "is odd")
		}
	}
}