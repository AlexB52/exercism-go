package isbn

import (
	"regexp"
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	matched, _ := regexp.Match(`\d{9}[\d|X]`, []byte(isbn))

	if len(isbn) != 10 || !matched {
		return false
	}

	var checksum int
	for i := 0; i < 10; i++ {
		n, err := strconv.Atoi(string(isbn[i]))
		if err != nil {
			n = 10
		}

		checksum += (10 - i) * n
	}

	return checksum%11 == 0
}
