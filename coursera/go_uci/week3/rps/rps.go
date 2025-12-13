package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var move rune
	machineMove := rand.Intn(3) // 0 = R, 1 = P, 2 = S
	fmt.Println("Choose either R P or S")
	fmt.Scanf("%c", &move)
	if move == 'R'{
		fmt.Printf("your move is %c\n", move)
	} else if move == 'P'{
		fmt.Printf("your move is %c\n", move)
	} else if move == 'S'{
		fmt.Printf("your move is %c\n", move)
	} else {
		fmt.Printf("illegal move %c\n", move)
		return
	}	

	fmt.Println("Machine move is" , machineMove)
	if move == 'R'  && machineMove == 1 {
		fmt.Printf("machine wins\n")
	} else if move == 'P' && machineMove == 2{
		fmt.Printf("machine wins\n")
	} else if move == 'S' && machineMove == 0{
		fmt.Printf("machine wins\n")
	} else {
		fmt.Printf("you drew or won\n")
	}

	
}