package cipher

import (
	"regexp"
	"strings"
	"unicode"
)

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.

type shift struct {
	Cipher
	offset int
}

type vigenere struct {
	Cipher
	key string
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance == 0 || distance/26 != 0 {
		return nil
	}
	return shift{offset: distance}
}

func (c shift) Encode(input string) string {
	return Code(c, input, encode())
}

func (c shift) Decode(input string) string {
	return Code(c, input, decode())
}

func (c shift) Distance(index int) int {
	return c.offset
}

func NewVigenere(key string) Cipher {
	re := regexp.MustCompile(`^[a-z]+[b-z]+$`)
	if len(key) == 0 || len(re.FindAllString(key, -1)) != 1 {
		return nil
	}
	return vigenere{key: key}
}

func (v vigenere) Encode(input string) string {
	re := regexp.MustCompile(`\w`)
	return Code(v, strings.Join(re.FindAllString(input, -1), ""), encode())
}

func (v vigenere) Decode(input string) string {
	return Code(v, input, decode())
}

func (v vigenere) Distance(index int) int {
	return int(v.key[index%len(v.key)] - 'a')
}

func Code(c Cipher, input string, transform func(rune, int) rune) string {
	var result []rune
	for i, r := range strings.ToLower(input) {
		if unicode.IsLetter(r) {
			result = append(result, transform(r, c.Distance(i)))
		}
	}
	return string(result)
}

func decode() func(rune, int) rune {
	return offset(-1)
}

func encode() func(rune, int) rune {
	return offset(1)
}

func offset(sign int) func(rune, int) rune {
	return func(r rune, distance int) rune {
		var utf8Distance = 'z' - 'a' + 1
		r += rune(sign * distance)
		if r > 'z' {
			r -= utf8Distance
		} else if r < 'a' {
			r += utf8Distance
		}
		return r
	}
}
