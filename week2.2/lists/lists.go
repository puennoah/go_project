package main

import (
	"fmt"
)

func main() {
	var head *ListElement
	var e, f ListElement
	xD := []int{4, 8, 16, 32, 64, 128}
	fmt.Println("Create list from", xD)
	head = &e
	e.data = 5
	e.next = &f
	f.data = 6
	fmt.Println("element e is:", e, "head is:", head)
	fmt.Println("element f is:", e.next, "head is", head)
	f.next = new(ListElement)
	f.next.data = 24
	fmt.Println("element f.next is:", f.next, "head is:", head)
	head.PrintList()
	FillList(xD, &head)
	fmt.Println("Head points to:", head)
	head.PrintList()
}

type ListElement struct {
	data int
	next *ListElement
}

//func createListElement(d int, ptr *ListElement) ListElement {
//	var element ListElement
//	element.data = d
//	element.next = ptr
//	return element
//}

func (h *ListElement) PrintList() {
	if h == nil {
		fmt.Println("###")
		return
	}
	fmt.Println(h.data, "->")
	h.next.PrintList()
}

func FillList(dataSlice []int, h **ListElement) {
	curEl := ListElement{dataSlice[0], nil}
	fmt.Println("curEl", curEl)
	*h = &curEl
	ptrCur := &curEl
	for i := 1; i < len(dataSlice); i++ {
		nextElem := ListElement{dataSlice[i], nil}
		fmt.Println("next", nextElem)
		ptrCur.next = &nextElem
		fmt.Println("curEl", curEl)
		ptrCur = &nextElem
	}
}
