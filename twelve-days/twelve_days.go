package twelve

import (
	"fmt"
	"strings"
)

var DayNumbers = []string{"", "first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
var Gifts = []string{
	"",
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

func Song() string {
	var song string
	for i := 1; i < 13; i++ {
		song += Verse(i)
		if i != 12 {
			song += "\n"
		}
	}
	return song
}

func Verse(i int) string {
	giftList := FormatGiftsList(GiftList(i))
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", DayNumbers[i], giftList)
}

func FormatGiftsList(gifts []string) string {
	var result string
	if len(gifts) == 1 {
		result = gifts[0]
	} else {
		result = fmt.Sprintf("%s, and %s", strings.Join(gifts[0:len(gifts)-1], ", "), gifts[len(gifts)-1])
	}
	return result
}

func GiftList(day int) []string {
	var result []string
	for i := day; i > 0; i-- {
		result = append(result, Gifts[i])
	}
	return result
}
