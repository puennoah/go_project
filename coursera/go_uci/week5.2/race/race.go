package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var wg sync.WaitGroup
	// Start 10 goroutines
	var mutex sync.Mutex
	fmt.Println("starting 10 threads each adding to count")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go incrementCount(&count, &wg, &mutex)
	}
	// Wait for all goroutines to finish
	wg.Wait()

	// Print the final count
	fmt.Println("Final count:", count)

}


func incrementCount(count *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()

	// Increment the count 1000 times
	for i := 0; i < 1000; i++ {
		(*mutex).Lock()
		(*count)++
		(*mutex).Unlock()
	}
}

