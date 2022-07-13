package yacht

func Score(dice []int, category string) int {
	var counts = map[int]int{}
	var scores = map[string]int{}
	var categories = []string{"", "ones", "twos", "threes", "fours", "fives", "sixes"}

	for _, d := range dice {
		counts[d]++
	}

	for n, c := range counts {
		scores[categories[n]] = n * c
		scores["choice"] += n * c

		if len(counts) == 2 && (c == 3 || c == 2) {
			scores["full house"] += n * c
		}

		if c >= 4 {
			scores["four of a kind"] = n * 4
		}
	}

	if len(counts) == 1 {
		scores["yacht"] = 50
	}

	if len(counts) == 5 && counts[1]+counts[2]+counts[3]+counts[4]+counts[5] == 5 {
		scores["little straight"] = 30
	}

	if len(counts) == 5 && counts[2]+counts[3]+counts[4]+counts[5]+counts[6] == 5 {
		scores["big straight"] = 30
	}

	return scores[category]
}
