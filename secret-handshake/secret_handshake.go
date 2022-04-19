package secret

var handshakeSigns = map[uint]string{
	1: "wink",
	2: "double blink",
	4: "close your eyes",
	8: "jump",
}

func Handshake(code uint) []string {
	var result []string

	if code&16 == 0 {
		for i := uint(1); i <= 8; i *= 2 {
			if code&i > 0 {
				result = append(result, handshakeSigns[i])
			}
		}
	} else {
		for i := uint(8); i >= 1; i /= 2 {
			if code&i > 0 {
				result = append(result, handshakeSigns[i])
			}
		}
	}

	return result
}
