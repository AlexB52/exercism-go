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
	verses := make([]string, 12)
	for i := 1; i < 13; i++ {
		verses[i-1] = Verse(i)
	}
	return strings.Join(verses, "\n")
}

func Verse(i int) string {
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", DayNumbers[i], FormatGiftsList(GiftList(i)))
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
