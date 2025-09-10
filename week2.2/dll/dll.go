package main

import (
	"fmt"
)

type Node struct {
	value rune
	prev  *Node
	next  *Node
}

type DLL struct {
	head *Node
	tail *Node
	size int
}

func main() {
	//list := NewDLL()

	// Helper function to build a list from a string
    buildList := func(s string) *DLL {
        list := NewDLL()
        for _, ch := range s {
            list.addToRear(ch)
        }
        return list
    }

    // Test cases
    tests := []struct {
        input      string
        ignoreCase bool
        expected   bool
    }{
        {"", true, true},         // empty string
        {"a", true, true},        // single char
        {"abba", true, true},     // even length palindrome
        {"aba", true, true},      // odd length palindrome
        {"abc", true, false},     // not a palindrome
        {"Abba", false, false},   // mismatch if case-sensitive
        {"Abba", true, true},     // works if ignoreCase = true
    }

    for _, t := range tests {
        list := buildList(t.input)
        result := list.isPalindrome(t.ignoreCase)
        fmt.Printf("Input: %-5s | ignoreCase: %-5v | Expected: %-5v | Got: %-5v\n",
            t.input, t.ignoreCase, t.expected, result)
    }
}

func NewDLL() *DLL {
	return &DLL{}
}

func checkInvariants(list *DLL) bool {
	var valid bool
	valid = true
	if list.size == 0 && (list.head != nil || list.tail != nil) {
		fmt.Println("Size 0 but list may not be empty")
		valid = false
	}

	if list.size > 0 && (list.head == nil || list.tail == nil) {
		fmt.Println("Size not 0 but list empty")
		valid = false
	}

	if list.head.prev != nil {
		fmt.Println("Head.prev is not nil")
		valid = false
	}
	if list.tail.next != nil {
		fmt.Println("Tail.next is not nil")
		valid = false
	}

	node := list.head
	for node != nil {
		if node.next != nil && node.next.prev != node {
			fmt.Println("Forward/backward link mismatch at node with value", node.value)
			valid = false
		}

		node = node.next
	}

	nodet := list.tail
	for nodet != nil {
		if nodet.prev != nil && nodet.prev.next != nodet {
			fmt.Println("Forward/backward link mismatch at node with value", nodet.value)
			valid = false
		}
		nodet = nodet.prev
	}

	counter := 0
	current := list.head
	if current != nil {
		counter++
		for current.next != nil {
			current = current.next
			counter++
		}
	}

	if counter != list.size {
		fmt.Println("Size mismatch: expect", list.size, "but count", counter)
		valid = false
	}
	return valid
}

func (list *DLL) addToFront(val rune) {
	newNode := &Node{value: val}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head.prev = newNode
		list.head = newNode
	}
	list.size++
	checkInvariants(list)
}

func (list *DLL) addToRear(val rune) {
	newNode := &Node{value: val}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}
	list.size++
	checkInvariants(list)
}

func (list *DLL) deleteFront() {
	if list.head != nil {
		list.head = list.head.next
		if list.head != nil {
			list.head.prev = nil
		}
	}
	if list.head == nil {
		list.tail = nil
	}
	if list.size > 0 {
		list.size--
	}
	checkInvariants(list)
}

func (list *DLL) deleteRear() {
	if list.tail != nil {
		list.tail = list.tail.prev
		if list.tail != nil {
			list.tail.next = nil
		}
	}
	if list.tail == nil {
		list.head = nil
	}
	if list.size > 0 {
		list.size--
	}
	checkInvariants(list)
}

func (list *DLL) delete(val rune) {
	node := list.head
	for node != nil {
		if node.value == val {
			if list.head == list.tail {
				list.head = nil
				list.tail = nil
			} else if node == list.head {
				list.head = list.head.next
				list.head.prev = nil
			} else if node == list.tail {
				list.tail = list.tail.prev
				list.tail.next = nil
			} else {
				node.next.prev = node.prev
				node.prev.next = node.next
			}
			list.size--
			break
		}
		node = node.next
	}
	checkInvariants(list)
}

func (list *DLL) findV(val rune) *Node {
	node := list.head
	for node != nil {
		if node.value == val {
			return node
		}
		node = node.next
	}
	return nil
}

func (list *DLL) isEmpty() bool {
	return list.size == 0 && list.head == nil && list.tail == nil
}

func (list *DLL) findLength() int {
	return list.size
}

func (list *DLL) insertPosition(i int, val rune) {
	newNode := &Node{value: val}
	if i <= 0 {
		list.addToFront(val)
	} else if i >= list.size {
		list.addToRear(val)
	} else {
		current := list.head
		for j := 0; j < i; j++ {
			current = current.next
		}
		newNode.prev = current
		newNode.next = current.next
		if current.next != nil {
		current.next.prev = newNode
		}
		current.next = newNode
		list.size++
	}
	checkInvariants(list)
}

func (list *DLL) deletePosition(i int) {
	if i <= 0 {
		list.deleteFront()
	} else if i >= list.size-1 {
		list.deleteRear()
	} else {
		current := list.head
		for j := 0; j < i; j++ {
			current = current.next
		}
		current.prev.next = current.next
		current.next.prev = current.prev
		list.size--
	}
	checkInvariants(list)
}

func (list *DLL) isPalindrome(ignoreCase bool) bool {
	if list.isEmpty() || list.size == 1 {
		return true
	}
	
    front := list.head
    rear := list.tail

    for front != nil && rear != nil && front != rear && front.prev != rear {
        fv := front.value
        rv := rear.value

        if ignoreCase {
            if fv >= 'A' && fv <= 'Z' {
                fv += 'a' - 'A'
            }
            if rv >= 'A' && rv <= 'Z' {
                rv += 'a' - 'A'
            }
        }

        if fv != rv {
            return false // mismatch found
        }

        front = front.next
        rear = rear.prev
    }

    return true
}

