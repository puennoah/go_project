package main

import (
	"fmt"
	"math/rand"
)

func main() {var move, machineMove int
	const (
		rock 	= 0
		paper	= 1
		scissors= 2
	)
	const ( 
		cRock	 = 'R'
		cPaper	 = 'P'
		cScissors= 'S'
	)
	var cMove rune
	var draws, wins, machineWins int
	var rounds int
	

	fmt.Println("How many rounds do you want to play?")
	fmt.Scanf("%d", &rounds)

	for i := 0; i < rounds; i++ {
		fmt.Println("\nRound", i + 1, ": Choose either R, P or S")
		fmt.Scanf("%c\n", &cMove)

		if cMove == cRock {
			move = rock
		} else if  cMove == cPaper {
			move = paper
		} else if cMove == cScissors {
			move = scissors
		} else {
			fmt.Println("-> illegal move")
			i--
			continue
		}
		

		machineMove = rand.Intn(3)
		if machineMove == rock {
			cMove = cRock
		} else if machineMove == paper{
			cMove = cPaper
		} else if machineMove == scissors {
			cMove = cScissors
		}
		fmt.Printf("machine plays %c\n", cMove)

		if move == machineMove {
			fmt.Println("-> draw")
			draws++
		} else if (move == paper) && (machineMove == scissors) {
			machineWins++
			fmt.Println("-> machine wins")
		} else if (move == rock) && (machineMove == paper) {
			machineWins++
			fmt.Println("-> machine wins")
		} else if (move == scissors) && (machineMove == rock) {
			machineWins++
			fmt.Println("-> machine wins")
		} else {
			wins++
			fmt.Println("-> you win")
		}
	}
	


}