package main

import "fmt"

func main() {
	fmt.Println("My Todo List")

	fmt.Println("Enter a todo: ")
	var input string
	fmt.Scanln(&input)
	fmt.Println("Todo:", input)
}
