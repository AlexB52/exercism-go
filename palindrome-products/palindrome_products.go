package palindrome

import (
	"errors"
	"strconv"
)

type Product struct {
	Value          int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
		return Product{}, Product{}, errors.New("fmin > fmax...")
	}

	var products = map[int]*Product{}
	var min, max *Product

	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			value := i * j
			if !IsPalindrome(value) {
				continue
			}

			p, ok := products[value]
			if !ok {
				products[value] = &Product{Value: value}
				p = products[value]
			}

			p.Factorizations = append(p.Factorizations, [2]int{i, j})

			if (min == nil) || (min.Value > value) {
				min = p
			}

			if (max == nil) || (max.Value < value) {
				max = p
			}
		}
	}

	if (min == nil) || (max == nil) {
		return Product{}, Product{}, errors.New("no palindromes...")
	}

	return *min, *max, nil
}

func IsPalindrome(value int) bool {
	s := []byte(strconv.Itoa(value))
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
