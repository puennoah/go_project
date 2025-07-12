package main

import "fmt"

var myInt int = 5

func main() {
	fmt.Println("Test simple syntax")
	fmt.Println("This is myInt", myInt)
	var myInt int = 6
	fmt.Println("This is inner myInt", myInt)
}
