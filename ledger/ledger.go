package ledger

import (
	"errors"
	"fmt"
	"sort"
	"time"
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

	var rows []string
	var s string
	if locale == "nl-NL" {
		s = fmt.Sprintf("%-10s | %-25s | %s\n", "Datum", "Omschrijving", "Verandering")
	} else if locale == "en-US" {
		s = fmt.Sprintf("%-10s | %-25s | %s\n", "Date", "Description", "Change")
	}
	rows = append(rows, s)
	for _, e := range entriesCopy {
		row, err := FormatRow(locale, currency, e)
		if err != nil {
			return "", err
		}
		rows = append(rows, row)
	}
	// Parallelism, always a great idea
	co := make(chan struct {
		i int
		s string
		e error
	})
	for i, et := range entriesCopy {
		go func(i int, entry Entry) {
			row, err := FormatRow(locale, currency, entry)
			if err != nil {
				co <- struct {
					i int
					s string
					e error
				}{e: err}
			}

			co <- struct {
				i int
				s string
				e error
			}{i: i, s: row}
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

func FormatRow(locale, currency string, entry Entry) (string, error) {
	t, err := time.Parse("2006-02-01", entry.Date)
	if err != nil {
		return "", errors.New("")
	}

	var description string
	if len(entry.Description) > 25 {
		description = fmt.Sprintf("%-22.22s...", entry.Description)
	} else {
		description = fmt.Sprintf("%-25s", entry.Description)
	}

	var symbol string
	if currency == "EUR" {
		symbol = "€"
	} else {
		symbol = "$"
	}

	var a, date string
	if locale == "nl-NL" {
		date = t.Format("01-02-2006")
		if entry.Change < 0 {
			a = fmt.Sprintf("%s %s-", symbol, FormatChange(entry.Change, ".", ","))
		} else {
			a = fmt.Sprintf("%s %s ", symbol, FormatChange(entry.Change, ".", ","))
		}
	} else if locale == "en-US" {
		date = t.Format("02/01/2006")
		if entry.Change < 0 {
			a = fmt.Sprintf("(%s%s)", symbol, FormatChange(entry.Change, ",", "."))
		} else {
			a = fmt.Sprintf(" %s%s ", symbol, FormatChange(entry.Change, ",", "."))
		}
	}

	return fmt.Sprintf("%10s | %s | %13s\n", date, description, a), nil
}

func FormatChange(change int, tsep, csep string) (result string) {
	if change < 0 {
		change *= -1
	}

	rest := fmt.Sprintf("%d", change/100)
	var parts []string
	for len(rest) > 3 {
		parts = append(parts, rest[len(rest)-3:])
		rest = rest[:len(rest)-3]
	}
	if len(rest) > 0 {
		parts = append(parts, rest)
	}
	for i := len(parts) - 1; i >= 0; i-- {
		result += parts[i] + tsep
	}
	result = result[:len(result)-1]

	return fmt.Sprintf("%s%s%02d", result, csep, change%100)
}
