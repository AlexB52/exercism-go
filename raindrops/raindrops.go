package raindrops

import "fmt"

func Convert(num int) string {
	result := ""

	if num%3 == 0 {
		result += "Pling"
	}

	if num%5 == 0 {
		result += "Plang"
	}

	if num%7 == 0 {
		result += "Plong"
	}

	if len(result) > 0 {
		return result
	} else {
		return fmt.Sprintf("%d", num)
	}
}
