package main

import (
	"luhn"
	"fmt"
)

func main() {
	fmt.Println(luhn.Valid("059"))
	// fmt.Println(luhn.Valid("8273 1232 7352 0569"))
	// fmt.Println(luhn.Valid("055 444 286"))
}
