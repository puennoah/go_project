package main

import "fmt"

func square(n int) int {
	return (n * n)
}

func main() {
	var n int
	fmt.Println("Input number to make it square value:")
	fmt.Scanf("%d", &n)
	fmt.Printf("The square of %d is %d.\n", n, square(n))
}
