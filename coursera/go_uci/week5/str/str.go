package main

import "fmt"

func main() {
	var c byte = 'a'
	var strOne string
	strOne = "My Idea?"
	fmt.Println("String Program")
	fmt.Println("First is:", strOne)
	fmt.Println("c is:", c)
	sum := len(strOne)
	for i := range sum {
		fmt.Printf("char %d is %c\n", i, strOne[i])
	}
}
