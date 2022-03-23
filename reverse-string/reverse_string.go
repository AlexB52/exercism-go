package reverse

func Reverse(input string) (result string) {
	runes, rInputs := []rune{}, []rune(input)

	for r := len(rInputs) - 1; r >= 0; r-- {
		runes = append(runes, rInputs[r])
	}

	return string(runes)
}
