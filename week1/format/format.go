package main

import "fmt"

func main() {
	i := 55
	x := 34567.89
	fmt.Println("Here is Println with defaults: i = ", i)
	fmt.Printf("Here is Printf with decimal format: i = %d\n", i)
	fmt.Printf("Here is Printf with  width 10d format: i = %10d\n", i)
	fmt.Printf("Here is Printf with binary format: i = %b\n", i)
	fmt.Printf("Here is Printf with o octal format: i = %o\n", i)
	fmt.Printf("Here is Printf with x hexadecimal format: i = %x\n", i)
	fmt.Printf("Here is Printf with O octal format: i = %O\n", i)
	fmt.Println("Floats")
	fmt.Println("Here is Println with defautls:", "x =", x)
	fmt.Printf("Here is Printf with f gormat: x = %f\n", x)
	fmt.Printf("Here is Printf with e format: x = %e\n", x)
	fmt.Printf("Here is Printf with 10g format: x = %10g\n", x)
	x = 1000 * x
	fmt.Printf("Here is Printf with g format: x = %g\n", x)
	fmt.Printf("Here is Printf with 10.2g format x = %10.2g\n", x)
}
