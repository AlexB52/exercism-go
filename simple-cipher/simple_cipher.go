package cipher

import (
	"regexp"
	"strings"
	"unicode"
)

type vigenere string

func NewCaesar() Cipher {
	return NewVigenere("d")
}

func NewShift(distance int) Cipher {
	if distance == 0 || distance/26 != 0 {
		return nil
	}
	key := string(transform('a', distance/distance, distance))
	return NewVigenere(key)
}

func NewVigenere(key string) Cipher {
	re := regexp.MustCompile(`^[b-z]$|^[a-z]+[b-z]+$`)
	if len(key) == 0 || len(re.FindAllString(key, -1)) != 1 {
		return nil
	}
	return vigenere(key)
}

func (v vigenere) Encode(input string) string {
	re := regexp.MustCompile(`\w`)
	input = strings.Join(re.FindAllString(input, -1), "")
	return v.Code(input, 1)
}

func (v vigenere) Decode(input string) string {
	return v.Code(input, -1)
}

func (v vigenere) Code(input string, sign int) string {
	var result []rune
	for i, r := range strings.ToLower(input) {
		if unicode.IsLetter(r) {
			distance := int(v[i%len(v)] - 'a')
			result = append(result, transform(r, sign, distance))
		}
	}
	return string(result)
}

func transform(r rune, sign int, distance int) rune {
	var utf8Distance = 'z' - 'a' + 1
	r += rune(sign * distance)
	if r > 'z' {
		r -= utf8Distance
	} else if r < 'a' {
		r += utf8Distance
	}
	return r
}
