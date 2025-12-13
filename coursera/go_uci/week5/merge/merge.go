/*Code for merge sort of randomly generated 10000 number*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("data of Int random numbers 0 to 10000")

	size := 10000
	dataSlice := make([]int, size)
	buffer := make([]int, size)
	for i := range dataSlice {
		dataSlice[i] = rand.Intn(10000)
	}

	fmt.Println("this is unorder data after randomly generated")
	fmt.Println(dataSlice[:50])
	mergeSort(dataSlice, buffer, 0, size-1)
	fmt.Println("\nthis is sorted data")
	fmt.Println(dataSlice[:50])

}

func mergeSort(slice, buffer []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	mergeSort(slice, buffer, left, mid)
	mergeSort(slice, buffer, mid+1, right)

	mergeSlice(slice, buffer, left, mid, right)
}

func mergeSlice(slice, buffer []int, left, mid, right int) {
	leftIndex := left
	rightIndex := mid + 1
	bufferIndex := left

	for leftIndex <= mid && rightIndex <= right {
		if slice[leftIndex] < slice[rightIndex] {
			buffer[bufferIndex] = slice[leftIndex]
			leftIndex++
			bufferIndex++
		} else {
			buffer[bufferIndex] = slice[rightIndex]
			rightIndex++
			bufferIndex++
		}
	}
	for leftIndex <= mid {
		buffer[bufferIndex] = slice[leftIndex]
		leftIndex++
		bufferIndex++
	}
	for rightIndex <= right {
		buffer[bufferIndex] = slice[rightIndex]
		rightIndex++
		bufferIndex++
	}

	for i := left; i <= right; i++ {
		slice[i] = buffer[i]
	}
}
