package rotationalcipher

func RotationalCipher(plain string, shiftKey int) (result string) {
	for _, r := range plain {
		var letter rune
		shiftLetter := rune(int(r) + shiftKey)

		switch {
		case 'a' <= r && r <= 'z':
			if shiftLetter > 'z' {
				shiftLetter -= 26
			}
			letter = shiftLetter
		case 'A' <= r && r <= 'Z':
			if shiftLetter > 'Z' {
				shiftLetter -= 26
			}
			letter = shiftLetter
		default:
			letter = r
		}

		result += string(letter)
	}
	return result
}
