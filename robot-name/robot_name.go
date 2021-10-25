package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

var namesUsed = make(map[string]bool)

type Robot struct {
	name string
}

func (robot *Robot) Reset() *Robot {
	robot.name = ""
	return robot
}

func (robot *Robot) Name() (string, error) {
	if robot.name == "" {
		robot.name = randomName()
	}

	return robot.name, nil
}

func randomName() string {
	rand.Seed(time.Now().UnixNano())
	name := ""
	name += randomLetter()
	name += randomLetter()
	name += randomNumber()

	if namesUsed[name] {
		name = randomName()
	} else {
		namesUsed[name] = true
	}

	return name
}

func randomLetter() string {
	return fmt.Sprintf("%c", 'A'+rune(rand.Intn(26)))
}

func randomNumber() string {
	return fmt.Sprintf("%03d", rand.Intn(1000))
}
