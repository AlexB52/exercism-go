package change

import (
	"errors"
	"math"
)

func Change(coins []int, target int) ([]int, error) {
	coinsPerChange := map[int][]int{0: []int{}}

	for i := 1; i <= target; i++ {
		var min = math.MaxUint8
		for _, coin := range coins {
			remainingChange, ok := coinsPerChange[i-coin]
			if ok && min > len(remainingChange) {
				min = len(remainingChange)
				coinsPerChange[i] = append([]int{coin}, remainingChange...)
			}
		}
	}

	result, ok := coinsPerChange[target]
	if !ok {
		return nil, errors.New("no change available")
	}
	return result, nil
}
