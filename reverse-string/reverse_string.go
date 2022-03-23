package reverse

func Reverse(input string) string {
	runes, inputRunes := []rune{}, []rune(input)

	for r := len(inputRunes) - 1; r >= 0; r-- {
		runes = append(runes, inputRunes[r])
	}

	return string(runes)
}
