package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var move, machineMove int
	const (
		rock     = 0
		paper    = 1
		scissors = 2
	)
	const (
		cRock     = 'R'
		cPaper    = 'P'
		cScissors = 'S'
	)
	//var cMove rune
	var draws, wins, machineWins int
	var rounds int
	const lengthForCheck int = 10
	recentMove := make([]int, 0, lengthForCheck)
	var rockCount, paperCount, scissorsCount int

	fmt.Println("How many rounds do you want to play?")
	fmt.Scanf("%d", &rounds)

	for i := 0; i < rounds; i++ {
		//fmt.Println("\nRound", i+1, ": Choose either R, P or S")
		//var input string
		//fmt.Scanln(&input)
		//if len(input) > 0 {
		//    cMove = rune(input[0])
		//} else {
		//    fmt.Println("Choose either R, P or S")
		//}
		move = rand.Intn(3)

		//if cMove == cRock {
		//	move = rock
		//} else if cMove == cPaper {
		//	move = paper
		//} else if cMove == cScissors {
		//	move = scissors
		//} else {
		//	fmt.Println("-> illegal move")
		//	i--
		//	continue
		//}
		recentMove = append(recentMove, move)
		if len(recentMove)  > lengthForCheck {
			recentMove = recentMove[1:]
		}
		// Check user move count
		rockCount, paperCount, scissorsCount = 0, 0, 0 
		for j := 0; j < len(recentMove); j++ {
			switch recentMove[j] {
			case 0 : rockCount++
			case 1 : paperCount++
			case 2 : scissorsCount++
			}
		}

		if rockCount > paperCount && rockCount > scissorsCount {
			machineMove = paper
		} else if paperCount > scissorsCount && paperCount > rockCount {
			machineMove = scissors
		} else if scissorsCount > rockCount && scissorsCount > paperCount {
			machineMove = rock
		} else {
			machineMove = rand.Intn(3)
		}

		//if machineMove == rock {
		//	cMove = cRock
		//} else if machineMove == paper{
		//	cMove = cPaper
		//} else if machineMove == scissors {
		//	cMove = cScissors
		//}
		//fmt.Printf("machine plays %c\n", cMove)

		switch move {
		case rock:
			if machineMove == rock {
				fmt.Println("-> draw")
				draws++
			} else if machineMove == paper {
				fmt.Println("-> machine wins")
				machineWins++
			} else {
				wins++
			}
		case paper:
			if machineMove == rock {
				fmt.Println("-> wins")
				wins++
			} else if machineMove == paper {
				fmt.Println("-> draws")
				draws++
			} else {
				machineWins++
			}
		case scissors:
			if machineMove == rock {
				fmt.Println("-> lose")
				machineWins++
			} else if machineMove == paper {
				fmt.Println("-> wins")
				wins++
			} else {
				draws++
			}
		}

	}
	fmt.Println("You win", wins, "times")
	fmt.Println("You lose,", machineWins, "times")
	fmt.Println("You draw", draws, "times")
}
