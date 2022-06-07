package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Structs & constants defined below FormatLedger

func FormatLedger(currency string, locale string, entries []Entry) (table string, err error) {
	currency_symbol, ok := CURRENCY_SYMBOL[currency]
	if !ok {
		return "", errors.New("unknown currency")
	}

	presenter, ok := LOCALE_PRESENTERS[locale]
	if !ok {
		return "", errors.New("unknown locale")
	}

	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)
	sort.Slice(entriesCopy, SortingEntriesAlgorithm(entriesCopy))

	table, err = BuildTable(
		presenter.Header,
		BuildRow(presenter.DateFormat, presenter.FormatChangeFunc(currency_symbol)),
		entriesCopy,
	)

	if err != nil {
		return "", err
	}

	return table, nil
}

var CURRENCY_SYMBOL = map[string]string{
	"EUR": "€",
	"USD": "$",
}

var LOCALE_PRESENTERS = map[string]LocalePresenter{
	"nl-NL": LocalePresenter{
		Header:     Row{"Datum", "Omschrijving", "Verandering"},
		DateFormat: "01-02-2006",
		FormatChangeFunc: func(symbol string) func(int) string {
			return func(change int) string {
				var format = "%s %s " // positive number format
				if change < 0 {
					format = "%s %s-"
				}
				return fmt.Sprintf(format, symbol, FormatChange(change, ".", ","))
			}
		},
	},
	"en-US": LocalePresenter{
		Header:     Row{"Date", "Description", "Change"},
		DateFormat: "02/01/2006",
		FormatChangeFunc: func(symbol string) func(int) string {
			return func(change int) string {
				var format = " %s%s " // positive number format
				if change < 0 {
					format = "(%s%s)"
				}
				return fmt.Sprintf(format, symbol, FormatChange(change, ",", "."))
			}
		},
	},
}

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type Row struct {
	Date, Description, Change string
}

type LocalePresenter struct {
	Header           Row
	DateFormat       string
	FormatChangeFunc func(string) func(int) string
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
	rows = append(rows, fmt.Sprintf("%-10s | %-25s | %s\n", header.Date, header.Description, header.Change))
	for _, entry := range entries {
		row, err := buildRow(entry)
		if err != nil {
			return "", err
		}
		rows = append(rows, fmt.Sprintf("%-10s | %-25s | %13s\n", row.Date, row.Description, row.Change))
	}

	return strings.Join(rows, ""), nil
}

func BuildRow(dateFormat string, formatChange func(int) string) func(e Entry) (Row, error) {
	return func(e Entry) (Row, error) {
		date, err := time.Parse("2006-02-01", e.Date)
		if err != nil {
			return Row{}, errors.New("invalid date")
		}

		var description string
		if len(e.Description) > 25 {
			description = fmt.Sprintf("%-22.22s...", e.Description)
		} else {
			description = fmt.Sprintf("%-25s", e.Description)
		}

		return Row{date.Format(dateFormat), description, formatChange(e.Change)}, nil
	}
}

func FormatChange(change int, tsep, csep string) (result string) {
	if change < 0 {
		change *= -1
	}

	rest := strconv.Itoa(change / 100)
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
