package scale

import (
	"strings"
)

type KeySignature uint8

const (
	sharp KeySignature = iota
	flat
)

var SIGNATURES = map[KeySignature][]string{
	sharp: []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"},
	flat:  []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"},
}

func Scale(tonic, interval string) (scale []string) {
	switch tonic {
	case "C", "a":
		scale = SIGNATURES[sharp]
	case "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#":
		scale = SIGNATURES[sharp]
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		scale = SIGNATURES[flat]
	}

	scale = scaleFromTonic(tonic, scale)
	scale = selectIntervals(scale, interval)
	return scale
}

func scaleFromTonic(tonic string, scale []string) (result []string) {
	for i, note := range scale {
		if strings.ToUpper(note) != strings.ToUpper(tonic) {
			continue
		}
		result = scale[i:]
		result = append(result, scale[:i]...)
		break
	}
	return result
}

func selectIntervals(scale []string, interval string) []string {
	if interval == "" {
		interval = "mmmmmmmmmmm"
	}

	var stepIndex int
	result := make([]string, len(interval)+1)
	result[0] = scale[0]
	for i, step := range interval {
		switch step {
		case 'M':
			stepIndex += 2
		case 'A':
			stepIndex += 3
		default:
			stepIndex += 1
		}
		result[i+1] = scale[stepIndex%12]
	}
	return result
}
