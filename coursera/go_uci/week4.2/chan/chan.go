package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Hello from goroutine 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Hello from goroutine 2"
	}()
	msg1 := <- ch
	fmt.Println(msg1)
	msg2 := <- ch
	fmt.Println(msg2)
}



