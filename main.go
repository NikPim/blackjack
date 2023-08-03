package main

import (
	"fmt"

	"github.com/NikPim/DeckOfCardAPI"
)
func Score(deck []DeckOfCardAPI.Card) int {
	sum := 0
	for _, val := range deck {
		sum += (Min(int(val.Rank), 10))
	}

	return sum
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	var input string
	var hand []DeckOfCardAPI.Card
	deck := DeckOfCardAPI.NewDeck()
	deck.Shuffle()
	for {
		fmt.Scanln(&input)
		if input == "s" {
        	break
		} else if input == "h" {
			hand, deck = append(hand, deck[0]), deck[1:]
			fmt.Println(Score(hand))
		} else {
			fmt.Println("s or h")
		}
    }
	fmt.Println(hand)
	fmt.Println(Score(hand))
}