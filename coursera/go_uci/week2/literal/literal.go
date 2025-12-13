package main

import "fmt"

func main() {
	var i, j, k int = 3, 4, 5
	fmt.Println("Explain Expressions")
	fmt.Println("i = ", i, "\tj = ", j, "\tk = ", k)
	i = j * k / 5
	fmt.Println("i = ", i, "\tj = ", j, "\tk = ", k)
	i = j * k + 99 % 6
	fmt.Println("i = ", i, "\tj = ", j, "\tk = ", k)
	i = j * (k + 99) % 6
	fmt.Println("i = ", i, "\tj = ", j, "\tk = ", k)
	i = j * ((k + 99) % 6)
	fmt.Println("i = ", i, "\tj = ", j, "\tk = ", k)
	i = -j * (k + 99)
	fmt.Println("i = ", i, "\tj = ", j, "\tk = ", k)
	i += -(j * (k + 99))
	fmt.Println("i = ", i, "\tj = ", j, "\tk = ", k)
	i = j * -k * 5
	fmt.Println("i = ", i, "\tj = ", j, "\tk = ", k)
}