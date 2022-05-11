package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {

	if locale != "nl-NL" && locale != "en-US" {
		return "", errors.New("")
	}

	if currency != "EUR" && currency != "USD" {
		return "", errors.New("")
	}

	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)

	sort.Slice(entriesCopy, func(i, j int) bool {
		if entriesCopy[i].Date < entriesCopy[j].Date {
			return true
		}

		if entriesCopy[i].Date > entriesCopy[j].Date {
			return false
		}

		if entriesCopy[i].Description < entriesCopy[j].Description {
			return true
		}

		if entriesCopy[i].Description > entriesCopy[j].Description {
			return false
		}

		return entriesCopy[i].Change < entriesCopy[j].Change
	})

	var s string
	if locale == "nl-NL" {
		s = fmt.Sprintf("%-10s | %-25s | %s\n", "Datum", "Omschrijving", "Verandering")
	} else if locale == "en-US" {
		s = fmt.Sprintf("%-10s | %-25s | %s\n", "Date", "Description", "Change")
	}

	// Parallelism, always a great idea
	co := make(chan struct {
		i int
		s string
		e error
	})
	for i, et := range entriesCopy {
		go func(i int, entry Entry) {
			d1, d2, d3, d4, d5 := entry.Date[0:4], entry.Date[4], entry.Date[5:7], entry.Date[7], entry.Date[8:10]
			if len(entry.Date) != 10 || d2 != '-' || d4 != '-' {
				co <- struct {
					i int
					s string
					e error
				}{e: errors.New("")}
			}
			var de string
			if len(entry.Description) > 25 {
				de = fmt.Sprintf("%-22.22s...", entry.Description)
			} else {
				de = fmt.Sprintf("%-25s", entry.Description)
			}
			var d string
			if locale == "nl-NL" {
				d = d5 + "-" + d3 + "-" + d1
			} else if locale == "en-US" {
				d = d3 + "/" + d5 + "/" + d1
			}
			negative := false
			cents := entry.Change
			if cents < 0 {
				cents = cents * -1
				negative = true
			}
			var a string
			if locale == "nl-NL" {
				if currency == "EUR" {
					a += "€"
				} else if currency == "USD" {
					a += "$"
				} else {
					co <- struct {
						i int
						s string
						e error
					}{e: errors.New("")}
				}
				a += " "
				centsStr := strconv.Itoa(cents)
				switch len(centsStr) {
				case 1:
					centsStr = "00" + centsStr
				case 2:
					centsStr = "0" + centsStr
				}
				rest := centsStr[:len(centsStr)-2]
				var parts []string
				for len(rest) > 3 {
					parts = append(parts, rest[len(rest)-3:])
					rest = rest[:len(rest)-3]
				}
				if len(rest) > 0 {
					parts = append(parts, rest)
				}
				for i := len(parts) - 1; i >= 0; i-- {
					a += parts[i] + "."
				}
				a = a[:len(a)-1]
				a += ","
				a += centsStr[len(centsStr)-2:]
				if negative {
					a += "-"
				} else {
					a += " "
				}
			} else if locale == "en-US" {
				if negative {
					a += "("
				}
				if currency == "EUR" {
					a += "€"
				} else if currency == "USD" {
					a += "$"
				} else {
					co <- struct {
						i int
						s string
						e error
					}{e: errors.New("")}
				}
				centsStr := strconv.Itoa(cents)
				switch len(centsStr) {
				case 1:
					centsStr = "00" + centsStr
				case 2:
					centsStr = "0" + centsStr
				}
				rest := centsStr[:len(centsStr)-2]
				var parts []string
				for len(rest) > 3 {
					parts = append(parts, rest[len(rest)-3:])
					rest = rest[:len(rest)-3]
				}
				if len(rest) > 0 {
					parts = append(parts, rest)
				}
				for i := len(parts) - 1; i >= 0; i-- {
					a += parts[i] + ","
				}
				a = a[:len(a)-1]
				a += "."
				a += centsStr[len(centsStr)-2:]
				if negative {
					a += ")"
				} else {
					a += " "
				}
			} else {
				co <- struct {
					i int
					s string
					e error
				}{e: errors.New("")}
			}
			var al int
			for range a {
				al++
			}
			co <- struct {
				i int
				s string
				e error
			}{i: i, s: d + strings.Repeat(" ", 10-len(d)) + " | " + de + " | " +
				strings.Repeat(" ", 13-al) + a + "\n"}
		}(i, et)
	}
	ss := make([]string, len(entriesCopy))
	for range entriesCopy {
		v := <-co
		if v.e != nil {
			return "", v.e
		}
		ss[v.i] = v.s
	}
	for i := 0; i < len(entriesCopy); i++ {
		s += ss[i]
	}
	return s, nil
}
