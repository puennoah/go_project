package main


type Stack[T any] struct {
	vals []T
}

func(s *Stack[T]) Push (val T) {
	s.vals = append(s.vals, val)
}