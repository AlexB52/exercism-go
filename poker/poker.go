package poker

import (
	"errors"
	"fmt"
	"regexp"
)

type Card string

func (c Card) Suit() string {
	return string([]rune(c)[1])
}

func (c Card) Value() string {
	return string([]rune(c)[0])
}

type Hand struct {
	Cards []Card
}

func (h Hand) Combination() string {
	suits, values := map[string]int{}, map[string]int{}
	for _, c := range h.Cards {
		suits[c.Suit()]++
		values[c.Value()]++
	}
	fmt.Println("suits", suits)
	fmt.Println("values", values)
	return ""
}

func BestHand(hands []string) ([]string, error) {
	re := regexp.MustCompile("\\b([2-9AJQK]|10)([♤♡♢♧])\\b")
	for _, h := range hands {
		cards := re.FindAllString(h, -1)
		if len(cards) != 5 {
			return nil, errors.New("invalid cards")
		}

		hand := Hand{}
		for _, c := range cards {
			hand.Cards = append(hand.Cards, Card(c))
		}
		hand.Combination()
		// fmt.Println("matches", matches)
		// fmt.Printf("%q\n", matches)
	}
	return nil, nil
}
