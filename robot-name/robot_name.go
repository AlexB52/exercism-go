package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

var existNames = make(map[string]bool)

type Robot struct {
	name string
}

func (robot *Robot) Reset() *Robot {
	robot.name = ""
	return robot
}

func (robot *Robot) Name() (string, error) {
	if robot.name == "" {
		robot.name = getRandomName()
	}

	return robot.name, nil
}

func getRandomName() string {
	rand.Seed(time.Now().UnixNano())
	name := ""
	name += randomLetter()
	name += randomLetter()
	name += randomNumber()
	name += randomNumber()
	name += randomNumber()

	if existNames[name] {
		name = getRandomName()
	} else {
		existNames[name] = true
	}

	return name
}

func randomLetter() string {
	return fmt.Sprintf("%c", 'A'+rune(rand.Intn(26)))
}

func randomNumber() string {
	return fmt.Sprintf("%d", rand.Intn(10))
}
