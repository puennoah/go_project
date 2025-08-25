//Program for test probability of straight flush hands
package main

import (
	"fmt"
	"math/rand"
)

type Suit int

const (
	club Suit = iota
	diamond
	heart
	spade
)

type Card struct {
	s   Suit
	pip int // Card number
}

func shuffle(d []Card) {
	for i := range d {
		j := rand.Intn(len(d))
		d[i], d[j] = d[j], d[i]
	}
}

func sortPip(c []Card) {
	for i := 0; i < len(c)-1; i++ {
		var swapped int
		for j := 0; j < (len(c) - 1 - i); j++ {
			if c[j].pip > c[j+1].pip {
				c[j], c[j+1] = c[j+1], c[j]
				swapped++
			}
		}
		if swapped == 0 {
			break
		}
	}
}

func (c Card) PrintCard() {
	var suitName, pipsName string
	switch c.s {
	case club:
		suitName = "Clubs"
	case diamond:
		suitName = "Diamonds"
	case heart:
		suitName = "Hearts"
	case spade:
		suitName = "Spades"
	}
	switch c.pip {
	case 1:
		pipsName = "Ace of "
	case 2:
		pipsName = "Two of "
	case 3:
		pipsName = "Three of "
	case 4:
		pipsName = "Four of "
	case 5:
		pipsName = "Five of "
	case 6:
		pipsName = "Six of "
	case 7:
		pipsName = "Seven of "
	case 8:
		pipsName = "Eight of "
	case 9:
		pipsName = "Nine of "
	case 10:
		pipsName = "Ten of "
	case 11:
		pipsName = "Jack of "
	case 12:
		pipsName = "Queen of "
	case 13:
		pipsName = "King of "
	}
	fmt.Printf("%s%s\n", pipsName, suitName)
}

func isStraightFlush(h []Card) bool {
	var ccount, dcount, hcount, scount int
	var clubs, diamonds, hearts, spades, sameSuits []Card
	for _, v := range h {
		switch v.s {
		case club:
			ccount++
			clubs = append(clubs, v)
		case diamond:
			dcount++
			diamonds = append(diamonds, v)
		case heart:
			hcount++
			hearts = append(hearts, v)
		case spade:
			scount++
			spades = append(spades, v)
		}
	}
	if ccount >= 5 {
		sameSuits = clubs
	} else if dcount >= 5 {
		sameSuits = diamonds
	} else if hcount >= 5 {
		sameSuits = hearts
	} else if scount >= 5 {
		sameSuits = spades
	} else {
		return false
	}
	var straightCounter int
	sortPip(sameSuits)
	if sameSuits[0].pip == 1 {
		sameSuits = append(sameSuits, Card{s: sameSuits[0].s, pip: 14})
	}
	for i := 0; i < (len(sameSuits) - 1); i++ {
		if sameSuits[i+1].pip-sameSuits[i].pip == 1 {
			straightCounter++
			if straightCounter == 4 {
				return true
			}
		} else if sameSuits[i+1].pip == sameSuits[i].pip {
			continue
		} else if sameSuits[i+1].pip-sameSuits[i].pip > 1 {
			straightCounter = 0
		}
	}
	return false
}

func main() {
	fmt.Println("Use Card Struct")
	deck := make([]Card, 52)
	var straightFlushCount int
	var totalCount int
	fmt.Println("Trials: ")
	fmt.Scanf("%d", &totalCount)
	for i := 0; i < len(deck); i++ {
		deck[i].pip = (i % 13) + 1
		deck[i].s = Suit(i / 13)
	}
	for i := 0; i < totalCount; i++ {
		handSize := 7
		numHands := len(deck) / handSize
		hand := make([][]Card, numHands)

		shuffle(deck)
		handIndex := 0

		for handIndex = 0; handIndex < numHands; handIndex++ {
			startIndex := handIndex * handSize
			endIndex := startIndex + handSize
			hand[handIndex] = deck[startIndex:endIndex]
			if isStraightFlush(hand[handIndex]) {
				straightFlushCount++
			}
		}
	}
	fmt.Println("Straight Flush count in 7 hands is :", straightFlushCount, "Per", totalCount, "plays")
	fmt.Printf("And the probability for straight flush is %f%%\n", float64(straightFlushCount)/float64(totalCount))
}
