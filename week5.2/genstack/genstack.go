package main

import (
	"fmt"
)


type Stack[T any] struct {
	vals []T
}

func(s *Stack[T]) Push (val T) {
	s.vals = append(s.vals, val)
}

// write pop func for exercise

func main() {
	fmt.Println("introduce generics")
	var intStack Stack[int]
	var stringStack Stack[string]
	fmt.Println(intStack)
	fmt.Println(stringStack)
	intStack.Push(15)
	intStack.Push(25)
	stringStack.Push("NPM")
	stringStack.Push("Aun")
	fmt.Println(intStack)
	fmt.Println(stringStack)
}