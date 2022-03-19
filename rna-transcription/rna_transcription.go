package strand

func ToRNA(dna string) (rna string) {
	DNAToRNA := map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}

	for _, n := range dna {
		rna += string(DNAToRNA[n])
	}

	return rna
}
