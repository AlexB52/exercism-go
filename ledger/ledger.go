package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"
)

var CURRENCY_SYMBOL = map[string]string{
	"EUR": "€",
	"USD": "$",
}

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type Row struct {
	date, description, change string
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

	var (
		header       Row
		dateFormat   string
		formatChange func(Entry) string
	)

	switch locale {
	case "nl-NL":
		header = Row{"Datum", "Omschrijving", "Verandering"}
		dateFormat = "01-02-2006"
		formatChange = FormatDutchChange(CURRENCY_SYMBOL[currency])
	case "en-US":
		header = Row{"Date", "Description", "Change"}
		formatChange = FormatUSChange(CURRENCY_SYMBOL[currency])
		dateFormat = "02/01/2006"
	}

	table, err = BuildTable(header, BuildRow(dateFormat, formatChange), entriesCopy)
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

func BuildTable(header Row, buildRow func(e Entry) (Row, error), entries []Entry) (result string, err error) {
	var rows []string
	rows = append(rows, fmt.Sprintf("%-10s | %-25s | %s\n", header.date, header.description, header.change))
	for _, entry := range entries {
		row, err := buildRow(entry)
		if err != nil {
			return "", err
		}
		rows = append(rows, fmt.Sprintf("%10s | %s | %13s\n", row.date, row.description, row.change))
	}

	return strings.Join(rows, ""), nil
}

func BuildRow(dateFormat string, formatChange func(Entry) string) func(e Entry) (Row, error) {
	return func(e Entry) (Row, error) {
		date, err := time.Parse("2006-02-01", e.Date)
		if err != nil {
			return Row{}, errors.New("")
		}

		var description string
		if len(e.Description) > 25 {
			description = fmt.Sprintf("%-22.22s...", e.Description)
		} else {
			description = fmt.Sprintf("%-25s", e.Description)
		}

		return Row{date.Format(dateFormat), description, formatChange(e)}, nil
	}
}

func FormatDutchChange(symbol string) func(e Entry) string {
	return func(e Entry) string {
		change := FormatChange(e.Change, ".", ",")
		if e.Change < 0 {
			return fmt.Sprintf("%s %s-", symbol, change)
		} else {
			return fmt.Sprintf("%s %s ", symbol, change)
		}
	}
}

func FormatUSChange(symbol string) func(e Entry) string {
	return func(e Entry) string {
		change := FormatChange(e.Change, ",", ".")
		if e.Change < 0 {
			return fmt.Sprintf("(%s%s)", symbol, change)
		} else {
			return fmt.Sprintf(" %s%s ", symbol, change)
		}
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
