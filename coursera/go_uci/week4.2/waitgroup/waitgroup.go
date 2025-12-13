package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func addSum() float64 {
	var data [10000000]float64
	var sum float64
	for i := 0; i < 10000000; i++ {
		data[i] = 10.0 * rand.Float64()
	}
	for _, d := range data {
		sum = sum + d
	}
	return sum
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5) // we will use 5 threads

	go func() {
		defer wg.Done()
		fmt.Println("sum is", addSum(), "first thread")
		fmt.Println(time.Now())
	}()

	go func() {
		defer wg.Done()
		fmt.Println("sum is", addSum(), "second thread")
		fmt.Println(time.Now())
	}()

	go func() {
		defer wg.Done()
		fmt.Println("sum is", addSum(), "third thread")
		fmt.Println(time.Now())
	}()

	go func() {
		defer wg.Done()
		fmt.Println("sum is", addSum(), "forth thread")
		fmt.Println(time.Now())
	}()

	go func() {
		defer wg.Done()
		fmt.Println("sum is", addSum(), "fifth thread")
		fmt.Println(time.Now())
	}()

	wg.Wait()
	fmt.Println(time.Now())


	for i := 0; i < 5; i++ { // just to compare with a normal loop
		fmt.Println(i + 1, addSum(), time.Now())
	}
}