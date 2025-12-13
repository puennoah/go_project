package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	
)

func main() {
	slice := make([]int, 0, 3)

	for {
		fmt.Println("Enter an integer or X to quit")
		var input string
		fmt.Scanln(&input)
		if strings.ToLower(input) == "x" {
			break
		}
		inputInt , err := strconv.Atoi(input)
		if err != nil {
			continue
		}
		slice = append(slice, inputInt)
		sort.Ints(slice)
		fmt.Println(slice)
	}
}