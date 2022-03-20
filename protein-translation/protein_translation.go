package protein

import (
	"errors"
)

var ErrStop = errors.New("ErrStop")
var ErrInvalidBase = errors.New("ErrInvalidBase")

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
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	default:
		return "", ErrInvalidBase
	}
}
