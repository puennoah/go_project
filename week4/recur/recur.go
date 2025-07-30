package main

import "fmt"

func main() {
	var n uint64
	fmt.Println("Enter n positive integer and compute n!")
	fmt.Scanf("%d\n", &n)
	fmt.Printf("n! is %d \n\n", factorial (n))
}

func factorial (n uint64) uint64 {
	if n <= 1 {
		return 1
	} else {
		return (n * factorial(n-1))
	}
}