package secret

var handshakeSigns = map[uint]string{
	1: "wink",
	2: "double blink",
	4: "close your eyes",
	8: "jump",
}

func Handshake(code uint) []string {
	var result []string
	for i, s := range handshakeSigns {
		if code&i > 0 {
			result = append(result, s)
		}
	}

	if code&16 > 0 {
		for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
			result[i], result[j] = result[j], result[i]
		}
	}

	return result
}
