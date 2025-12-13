package main

import (
	"reflect"
	"testing"
)

// Test Push and isEmpty
func TestPushAndIsEmpty(t *testing.T) {
	var s Stack[int]
	if !s.isEmpty() {
		t.Errorf("expected stack to be empty initially")
	}

	s.Push(10)
	if s.isEmpty() {
		t.Errorf("expected stack to have elements after push")
	}
	if s.Top() != 10 {
		t.Errorf("expected top to be 10, got %v", s.Top())
	}
}

// Test Pop
func TestPop(t *testing.T) {
	var s Stack[int]
	s.Push(5)
	s.Push(15)

	val, ok := s.Pop()
	if !ok || val != 15 {
		t.Errorf("expected Pop to return 15, got %v", val)
	}

	val, ok = s.Pop()
	if !ok || val != 5 {
		t.Errorf("expected Pop to return 5, got %v", val)
	}

	_, ok = s.Pop()
	if ok {
		t.Errorf("expected Pop on empty stack to return ok=false")
	}
}

// Test Top
func TestTop(t *testing.T) {
	var s Stack[string]
	s.Push("hello")
	if s.Top() != "hello" {
		t.Errorf("expected Top to return 'hello', got %v", s.Top())
	}
}

// Test CopyFromSlice
func TestCopyFromSlice(t *testing.T) {
	var s Stack[int]
	s.CopyFromSlice([]int{1, 2, 3})
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(s.vals, expected) {
		t.Errorf("expected stack vals %v, got %v", expected, s.vals)
	}
}

// Test CopyToSlice
func TestCopyToSlice(t *testing.T) {
	var s Stack[int]
	s.CopyFromSlice([]int{4, 5, 6})
	copied := s.CopyToSlice()
	expected := []int{4, 5, 6}
	if !reflect.DeepEqual(copied, expected) {
		t.Errorf("expected copied slice %v, got %v", expected, copied)
	}

	// ensure stack not modified
	if !reflect.DeepEqual(s.vals, expected) {
		t.Errorf("expected stack vals to remain %v, got %v", expected, s.vals)
	}
}
