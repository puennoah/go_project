package main

import (
	"fmt"
	"strings"
)

func main() {
	var strInput string
	fmt.Println("Palindrome Program - Input string")
	fmt.Scanf("%s", &strInput)
	sum := len(strInput)
	strTest := strings.ToLower(strInput)
	for i := range sum {
		fmt.Printf("char %d is %c\n", i, strTest[i])
		if strTest[i] == strTest[sum-(i+1)] {
			continue
		} else {
			fmt.Println("Not a palindrome")
			return
		}
	}
	fmt.Println("Is a palindrome")
}
