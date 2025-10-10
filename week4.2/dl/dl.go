package main

import (
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		mess := "In anonymous function"
		c1 <- mess
	stress := <- c2
	fmt.Println(mess, "and", stress)
	}()
	mess := "In main() message"
	stress := <- c1
	c2 <- mess
	fmt.Println(mess, "and", stress)
}