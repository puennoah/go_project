// code for dining philosophers that each philosopher get the eat 3 times
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numPhilosophers = 5
	numForks        = 5
	numLadle        = 1
)

type Philosopher struct {
	id        int
	leftFork  *sync.Mutex
	rightFork *sync.Mutex
	ladle     *sync.Mutex
}

func (p *Philosopher) eat(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		think := rand.Intn(2) + 1
		fmt.Printf("Philosopher %d is thinking for %d seconds\n", p.id, think)
		time.Sleep(time.Duration(think) * time.Second)

		scoop := rand.Intn(2) + 1
		p.ladle.Lock()
		fmt.Printf("Philosopher %d picked up the ladle\n", p.id)

		fmt.Printf("Philosopher %d is scooping for %d seconds\n", p.id, scoop)
		time.Sleep(time.Duration(scoop) * time.Second)
		p.ladle.Unlock()
		fmt.Printf("Philosopher %d put down the ladle\n", p.id)

		if p.id%2 == 0 {
			p.leftFork.Lock()
			fmt.Printf("Philosopher %d picked up the left fork\n", p.id)
			p.rightFork.Lock()
			fmt.Printf("Philosopher %d picked up the right fork\n", p.id)
		} else {
			p.rightFork.Lock()
			fmt.Printf("Philosopher %d picked up the right fork\n", p.id)
			p.leftFork.Lock()
			fmt.Printf("Philosopher %d picked up the left fork\n", p.id)
		}

		eat := rand.Intn(2) + 1
		fmt.Printf("Philosopher %d is eating for %d seconds\n", p.id, eat)
		time.Sleep(time.Duration(eat) * time.Second)

		p.rightFork.Unlock()
		fmt.Printf("Philosopher %d put down the right fork\n", p.id)
		p.leftFork.Unlock()
		fmt.Printf("Philosopher %d put down the left fork\n", p.id)
	}
}

func main() {
	forks := make([]*sync.Mutex, numForks)
	ladle := &sync.Mutex{}
	for i := range forks {
		forks[i] = &sync.Mutex{}
	}

	philosophers := make([]*Philosopher, numPhilosophers)
	for i := range philosophers {
		leftFork := forks[i]
		rightFork := forks[(i+1)%numForks]
		philosophers[i] = &Philosopher{id: i, leftFork: leftFork, rightFork: rightFork, ladle: ladle}
	}
	var wg sync.WaitGroup
	wg.Add(numPhilosophers) // remember 5 philosophers - each is a thread
	for _, p := range philosophers {
		go p.eat(&wg)
	}
	wg.Wait()
}
