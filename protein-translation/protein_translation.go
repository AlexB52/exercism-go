package protein

import (
	"errors"
)

var ErrStop = errors.New("ErrStop")
var ErrInvalidBase = errors.New("ErrInvalidBase")

var RNATranslations = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGG": "Tryptophan",
}

func FromRNA(rna string) ([]string, error) {
	var result []string
translation:
	for i := 0; i < len(rna); i += 3 {
		switch t, err := FromCodon(rna[i : i+3]); err {
		case ErrInvalidBase:
			return []string{}, err
		case ErrStop:
			break translation
		default:
			result = append(result, t)
		}
	}
	return result, nil
}

func FromCodon(codon string) (string, error) {
	switch codon {
	case "UAG", "UGA", "UAA":
		return "", ErrStop
	default:
		if t, ok := RNATranslations[codon]; ok {
			return t, nil
		} else {
			return "", ErrInvalidBase
		}
	}
}
