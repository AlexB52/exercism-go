package dna

import "errors"

type Histogram map[rune]int

type DNA []rune

func (d DNA) Counts() (Histogram, error) {
	var h = Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for _, nucleotide := range d {
		_, ok := h[nucleotide]

		if !ok {
			return nil, errors.New("Invalid nucleotide")
		}

		h[nucleotide]++
	}

	return h, nil
}
