package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

var namesUsed = make(map[string]bool)

var maxNamesCount int = 26 * 26 * 10 * 10 * 10 // 676,000
var randomNames = generateNames()
var nameIndex = 0

type Robot struct {
	name string
}

func (robot *Robot) Reset() *Robot {
	robot.name = ""
	return robot
}

func (robot *Robot) Name() (string, error) {
	if robot.name == "" {
		name, error := randomName()
		if error != nil {
			return "", error
		}
		robot.name = name
	}

	return robot.name, nil
}

func randomName() (string, error) {
	if nameIndex >= maxNamesCount {
		return "", fmt.Errorf("no more names available")
	}

	result := randomNames[nameIndex]
	nameIndex++
	return result, nil
}

func generateNames() []string {
	var names = make([]string, maxNamesCount)
	var position int

	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			for n := 0; n < 1000; n++ {
				var name string
				name += fmt.Sprintf("%c", 'A'+i)
				name += fmt.Sprintf("%c", 'A'+j)
				name += fmt.Sprintf("%03d", n)
				names[position] = name
				position++
			}
		}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(maxNamesCount, func(i, j int) { names[i], names[j] = names[j], names[i] })
	return names
}
