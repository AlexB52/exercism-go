package queenattack

import (
	"fmt"
	"regexp"
	"strconv"
)

type Position struct {
	x int
	y int
}

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	var LetterMap = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8}

	re := regexp.MustCompile("([a-h])([1-8])")
	p1Matches := re.FindStringSubmatch(whitePosition)
	p2Matches := re.FindStringSubmatch(blackPosition)

	if whitePosition == blackPosition || p1Matches == nil || p2Matches == nil {
		return false, fmt.Errorf("invalid position")
	}

	y1, _ := strconv.Atoi(p1Matches[2])
	y2, _ := strconv.Atoi(p2Matches[2])

	p1 := Position{LetterMap[p1Matches[1]], y1}
	p2 := Position{LetterMap[p2Matches[1]], y2}

	sameLine := p1.x == p2.x
	sameColumn := p1.y == p2.y
	sameDiagonal2 := p1.x-p2.x == p1.y-p2.y || p1.x-p2.x == p2.y-p1.y

	if sameLine || sameColumn || sameDiagonal2 {
		return true, nil
	}

	return false, nil
}
