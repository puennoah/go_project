package main

import "fmt"

func main() {
	fmt.Println("Use Pointers in GO")
	i := 55
	pointerToi := &i
	fmt.Println("i =", i)
	fmt.Println("address of i =", pointerToi)
	*pointerToi++
	fmt.Println("i =", i)
	j := 45
	pointerToi = &j
	fmt.Println("address of j =", pointerToi)
	pointerToi = nil
	fmt.Println("value of pointerToi =", pointerToi)
}
