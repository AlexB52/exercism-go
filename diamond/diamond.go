package diamond

import (
	"errors"
	"strings"
)

func Gen(char byte) (string, error) {
	if char < 'A' || 'Z' < char {
		return "", errors.New("out of range")
	}
	l1 := int(char - 'A')
	l2 := 2*l1 + 1

	var result = make([]string, l2)

	for i := 0; i <= l1; i++ {
		letter := char - byte(i)
		result[l1-i] = GenLine(letter, i, l2)
		result[l1+i] = GenLine(letter, i, l2)
	}

	return strings.Join(result, "\n"), nil
}

func GenLine(char byte, pos, length int) string {
	var result = make([]byte, length)
	for i := range result {
		result[i] = '·'
	}
	result[pos] = char
	result[length-pos-1] = char
	return string(result)
}

// "...A..."
// "..B.B.."
// ".C...C."
// "D.....D"
// ".C...C."
// "..B.B.."
// "...A..."

// .....A.....
// ....B.B....
// ...C...C...
// ..D.....D..
// .E.......E.
// F.........F
// .E.......E.
// ..D.....D..
// ...C...C...
// ....B.B....
// .....A.....
