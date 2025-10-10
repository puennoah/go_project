package main

import (
	"fmt"
	"math"
)

type SqrtResult struct {
	num int
	sqrt float64
}

func worker(id int, jobs <- chan int, results chan <- SqrtResult) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		results <- SqrtResult{num: j, sqrt: math.Sqrt(float64(j))}
		fmt.Println("Worker", id, "finish job", j)
	}
}

func main() {
	const numJobs = 10
	jobs := make(chan int, numJobs)
	results := make(chan SqrtResult, numJobs)

	const numWorkers = 10
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	for j := 1;  j <= numJobs; j++ {
		jobs <- j
	}
	for j := 1; j <= numJobs; j++ {
		res := <- results
		fmt.Println("sqrt of", res.num, "is", res.sqrt)
	}
	close(jobs)
}