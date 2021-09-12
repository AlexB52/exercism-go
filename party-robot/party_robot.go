package partyrobot

import "fmt"

func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

func AssignTable(name string, table int, neighbour string, direction string, distance float64) string {
	result := fmt.Sprintf("%s\n", Welcome(name))
	result += fmt.Sprintf("You have been assigned to table %X. Your table is %s, exactly %.1f meters from here.\n", table, direction, distance)
	result += fmt.Sprintf("You will be sitting next to %s.", neighbour)
	return result
}
