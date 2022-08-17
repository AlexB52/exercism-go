package queenattack

import (
	"fmt"
	"regexp"
)

func CanQueenAttack(whiteQ, blackQ string) (bool, error) {
	re := regexp.MustCompile("([a-h])([1-8])")
	if whiteQ == blackQ || !re.Match([]byte(whiteQ)) || !re.Match([]byte(blackQ)) {
		return false, fmt.Errorf("invalid position")
	}

	xDiff, yDiff := whiteQ[0]-blackQ[0], whiteQ[1]-blackQ[1]

	return xDiff == 0 || yDiff == 0 || xDiff == yDiff || xDiff == -yDiff, nil
}
