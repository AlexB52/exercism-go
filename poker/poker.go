package poker

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

type HandRank int
type Card string

const (
	HighCard HandRank = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

func (c Card) Suit() string {
	return string([]rune(c)[1])
}

func (c Card) Value() string {
	return string([]rune(c)[0])
}

type Hand struct {
	Cards  []Card
	Suits  map[string]int
	Values map[string]int
	List   []int
}

func (h *Hand) Score() HandRank {
	switch {
	default:
		return HighCard
	}
}

func BuildHand(cards []string) Hand {
	hand := Hand{
		Cards:  make([]Card, 5),
		Suits:  map[string]int{},
		Values: map[string]int{},
		List:   []int{},
	}

	for i, c := range cards {
		card := Card(c)
		hand.Cards[i] = card
		hand.Suits[card.Suit()]++
		hand.Values[card.Value()]++
		i, _ := strconv.Atoi(card.Value())
		hand.List = append(hand.List, i)
	}

	sort.Ints(hand.List)

	fmt.Println("suits", hand.Suits)
	fmt.Println("values", hand.Values)
	fmt.Println("list", hand.List)
	fmt.Println()

	return hand
}

func BestHand(dealtHands []string) ([]string, error) {
	cardRegex := "([2-9AJQK]|10)([♤♡♢♧])"

	reValidation := regexp.MustCompile(fmt.Sprintf("^(%s\\s){4}%s$", cardRegex, cardRegex))
	for _, h := range dealtHands {
		if !reValidation.Match([]byte(h)) {
			return nil, errors.New("invalid cards")
		}
	}

	var hands = []Hand{}
	var highestScore HandRank
	for _, h := range dealtHands {
		cards := regexp.MustCompile(cardRegex).FindAllString(h, -1)
		fmt.Printf("%q\n ", cards)
		hands = append(hands, BuildHand(cards))
		// fmt.Println("matches", matches)
		// fmt.Printf("%q\n", matches)
	}
	return nil, nil
}
