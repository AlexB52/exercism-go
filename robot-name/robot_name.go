package robotname

import (
	"fmt"
	"time"
	"math/rand"
)

var names = map[string]bool{}

type Robot struct {
	name string
}

func New() *Robot {
	return &Robot{getRandomName()}
}

func getRandomName() string {
	rand.Seed(time.Now().UnixNano())
	name := ""
	name += randomLetter()
	name += randomLetter()
	name += randomNumber()
	name += randomNumber()
	name += randomNumber()

	if !names[name] {
		names[name] = true
		return name
	} else {
		return getRandomName()
	}
}

func randomLetter() string {
	return fmt.Sprintf("%c", 'A' + rune(rand.Intn(26)))
}

func randomNumber() string {
	return fmt.Sprintf("%d", rand.Intn(10))
}

func (robot Robot) Reset() Robot {
	robot.name = getRandomName()
	return robot
}

func (robot Robot) Name() (string, error){
	return robot.name, nil
}