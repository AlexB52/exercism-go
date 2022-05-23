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

type Row struct {
	description, amount, date string
}

func FormatLedger(currency string, locale string, entries []Entry) (table string, err error) {
	if locale != "nl-NL" && locale != "en-US" {
		return "", errors.New("")
	}

	if currency != "EUR" && currency != "USD" {
		return "", errors.New("")
	}

	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)

	sort.Slice(entriesCopy, SortingEntriesAlgorithm(entriesCopy))

	switch locale {
	case "nl-NL":
		table, err = BuildTable(Row{"Omschrijving", "Verandering", "Datum"}, entriesCopy, currency)
	case "en-US":
		table, err = BuildUSTable(entriesCopy, currency)
	}

	if err != nil {
		return "", err
	}

	return table, nil
}

func SortingEntriesAlgorithm(entriesCopy []Entry) func(i, j int) bool {
	return func(i, j int) bool {
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
	}
}

func BuildTable(header Row, entries []Entry, currency string) (result string, err error) {
	var rows []string
	rows = append(rows, fmt.Sprintf("%-10s | %-25s | %s\n", header.date, header.description, header.amount))
	for _, entry := range entries {
		row, err := BuildDutchRow(entry, currency)
		if err != nil {
			return "", err
		}
		rows = append(rows, fmt.Sprintf("%10s | %s | %13s\n", row.date, row.description, row.amount))
	}

	return strings.Join(rows, ""), nil
}

func BuildDutchTable(entries []Entry, currency string) (result string, err error) {
	var rows []string
	rows = append(rows, fmt.Sprintf("%-10s | %-25s | %s\n", "Datum", "Omschrijving", "Verandering"))
	for _, entry := range entries {
		row, err := BuildDutchRow(entry, currency)
		if err != nil {
			return "", err
		}
		rows = append(rows, fmt.Sprintf("%10s | %s | %13s\n", row.date, row.description, row.amount))
	}

	return strings.Join(rows, ""), nil
}

func BuildDutchRow(e Entry, currency string) (Row, error) {
	date, err := time.Parse("2006-02-01", e.Date)
	if err != nil {
		return Row{}, errors.New("")
	}

	var amount string
	if e.Change < 0 {
		amount = fmt.Sprintf("%s %s-", Symbol(currency), FormatChange(e.Change, ".", ","))
	} else {
		amount = fmt.Sprintf("%s %s ", Symbol(currency), FormatChange(e.Change, ".", ","))
	}

	return Row{FormatDescription(e.Description), amount, date.Format("01-02-2006")}, nil
}

func BuildUSTable(entries []Entry, currency string) (result string, err error) {
	var rows []string
	rows = append(rows, fmt.Sprintf("%-10s | %-25s | %s\n", "Date", "Description", "Change"))
	for _, entry := range entries {
		row, err := BuildUSRow(entry, currency)
		if err != nil {
			return "", err
		}
		rows = append(rows, fmt.Sprintf("%10s | %s | %13s\n", row.date, row.description, row.amount))
	}

	return strings.Join(rows, ""), nil
}

func BuildUSRow(e Entry, currency string) (Row, error) {
	date, err := time.Parse("2006-02-01", e.Date)
	if err != nil {
		return Row{}, errors.New("")
	}

	var amount string
	if e.Change < 0 {
		amount = fmt.Sprintf("(%s%s)", Symbol(currency), FormatChange(e.Change, ",", "."))
	} else {
		amount = fmt.Sprintf(" %s%s ", Symbol(currency), FormatChange(e.Change, ",", "."))
	}

	return Row{FormatDescription(e.Description), amount, date.Format("02/01/2006")}, nil
}

func Symbol(currency string) string {
	if currency == "EUR" {
		return "€"
	} else {
		return "$"
	}
}

func FormatDescription(s string) string {
	if len(s) > 25 {
		return fmt.Sprintf("%-22.22s...", s)
	} else {
		return fmt.Sprintf("%-25s", s)
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
