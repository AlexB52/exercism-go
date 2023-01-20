package scale

import (
	"strings"
)

type KeySignature uint8

const (
	Sharp KeySignature = iota
	Flat
)

var STEP_INTERVALS = map[rune]int{'m': 1, 'M': 2, 'A': 3}
var SIGNATURES = map[KeySignature][]string{
	Sharp: []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"},
	Flat:  []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"},
}

func Scale(tonic, interval string) []string {
	if interval == "" {
		interval = "mmmmmmmmmmm"
	}

	signature := Signature(tonic)
	stepIndex := TonicIndex(tonic, signature)

	scale := make([]string, len(interval)+1)
	scale[0] = signature[stepIndex]
	for i, step := range interval {
		stepIndex += STEP_INTERVALS[step]
		scale[i+1] = signature[stepIndex%len(signature)]
	}

	return scale
}

func Signature(tonic string) (signature []string) {
	switch tonic {
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		return SIGNATURES[Flat]
	default:
		return SIGNATURES[Sharp]
	}
}

func TonicIndex(tonic string, scale []string) int {
	for i, note := range scale {
		if strings.ToUpper(note) == strings.ToUpper(tonic) {
			return i
		}
	}
	return -1
}
