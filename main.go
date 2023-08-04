package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/NikPim/DeckOfCardAPI"
)

type Hand []DeckOfCardAPI.Card

func (hand Hand) String() string {
    var strs []string
    for _, card := range hand {
        if card.Suit == DeckOfCardAPI.Joker{
			strs = append(strs, card.Suit.String()) 
		} else{
			strs = append(strs, fmt.Sprintf("%s of %ss", card.Rank.String(), card.Suit.String())) 
		}
    }
    return strings.Join(strs, ", ")
}

func Score(hand Hand) int {
	sum := 0
	firstAce := true
	for _, card := range hand {
		if card.Rank == DeckOfCardAPI.Ace && firstAce {
			sum += 11
			firstAce = false
		} else {
			sum += (Min(int(card.Rank), 10))
		}
	}
	if sum > 21 && !firstAce {
		sum -= 10
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

func cls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func PrintResult(hand Hand, final bool) {
	cls()
	entity := "current"
	if final {
		entity = "final"
	}
	fmt.Printf("\nYour %s hand is: %s", entity, hand)
	fmt.Printf("\nYour %s score is: %d\n", entity, Score(hand))
}

func drawCard(hand *Hand, deck *DeckOfCardAPI.Deck) {
	*hand = append(*hand, (*deck)[0])
	*deck = (*deck)[1:]
}

func Blackjack() {
	var input string
	var hand Hand
	deck := DeckOfCardAPI.NewDeck(true)
	for i:= 0; i < 2;i++ {
		drawCard(&hand, &deck)
	}
	PrintResult(hand, false)
	for input != "s"{
		fmt.Println("Do you want to (h)it or to (s)tand?")
		fmt.Scanln(&input)
		if input == "h" {
			drawCard(&hand, &deck)
			PrintResult(hand, false)
		}
    }
	PrintResult(hand, true)
}

func main() {
	Blackjack()
}