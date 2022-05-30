package brackets

var PAIRS = map[rune]rune{
	'}': '{',
	']': '[',
	')': '(',
}

func Bracket(input string) bool {
	var inputs []rune
	for i := 0; i < len(input); i++ {
		current := rune(input[i])
		switch current {
		case '[', '{', '(':
			inputs = append(inputs, current)
		case ']', '}', ')':
			if len(inputs) > 0 && inputs[len(inputs)-1] == PAIRS[current] {
				inputs = inputs[:len(inputs)-1]
			} else {
				return false
			}
		default:
		}
	}

	return len(inputs) == 0
}
