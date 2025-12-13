package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s", &input)
	input = strings.ToLower(input)
	if len(input) < 2 {
		fmt.Println("Not Found!")
		return
	}

	if input[0] == 'i' && strings.Contains(input, "a") && input[len(input) - 1] == 'n' {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}