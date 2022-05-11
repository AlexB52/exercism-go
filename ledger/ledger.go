package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strings"
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
	return strings.Join(rows, ""), nil
}

func FormatDutchRow(currency string, entry Entry) (string, error) {
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

	var a string
	if entry.Change < 0 {
		a = fmt.Sprintf("%s %s-", symbol, FormatChange(entry.Change, ".", ","))
	} else {
		a = fmt.Sprintf("%s %s ", symbol, FormatChange(entry.Change, ".", ","))
	}

	return fmt.Sprintf("%10s | %s | %13s\n", t.Format("01-02-2006"), description, a), nil
}

func FormatUSRow(currency string, entry Entry) (string, error) {
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

	var a string
	if entry.Change < 0 {
		a = fmt.Sprintf("(%s%s)", symbol, FormatChange(entry.Change, ",", "."))
	} else {
		a = fmt.Sprintf(" %s%s ", symbol, FormatChange(entry.Change, ",", "."))
	}

	return fmt.Sprintf("%10s | %s | %13s\n", t.Format("02/01/2006"), description, a), nil
}

func FormatRow(locale, currency string, entry Entry) (result string, err error) {
	if locale == "nl-NL" {
		result, err = FormatDutchRow(currency, entry)
	} else if locale == "en-US" {
		result, err = FormatUSRow(currency, entry)
	}
	return result, err
}

func FromatDescription(entry Entry) string {
	if len(entry.Description) > 25 {
		return fmt.Sprintf("%-22.22s...", entry.Description)
	} else {
		return fmt.Sprintf("%-25s", entry.Description)
	}
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
