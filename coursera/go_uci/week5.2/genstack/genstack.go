// Generic stack and testing
package main

import (
	"fmt"
	"slices"
)

type Stack[T any] struct {
	vals []T
}

func main() {
	fmt.Println("introduce generics")
	slice := []int{1, 2, 20}
	var intStack Stack[int]
	fmt.Println(intStack)
	intStack.Push(15)
	intStack.Push(25)
	fmt.Println(intStack)
	intStack.Pop()
	fmt.Println(intStack)
	intStack.CopyFromSlice(slice)
	fmt.Println("This is slice of int:", slice)
	fmt.Println("This is stack after CopyFromSlice", intStack)
	fmt.Println("is intStack empty :", intStack.isEmpty())
	sliceCopied := intStack.CopyToSlice()
	fmt.Println("Here is the copied slice:", sliceCopied)

}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}

	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	fmt.Println(top, "Popped")
	return top, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.vals) == 0
}

func (s *Stack[T]) Top() T {
	return s.vals[len(s.vals)-1]
}

func (s *Stack[T]) CopyFromSlice(slice []T) {
	s.vals = append(s.vals, slice...)
}

func (s *Stack[T]) CopyToSlice() []T {
	return slices.Clone(s.vals)
}
