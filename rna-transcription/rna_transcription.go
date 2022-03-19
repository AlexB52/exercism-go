package strand

func ToRNA(dna string) (newDNA string) {
	transcription := map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}

	for _, nucleotide := range dna {
		newDNA += string(transcription[nucleotide])
	}

	return newDNA
}
