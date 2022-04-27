package variablelengthquantity

import (
	"errors"
	"math"
)

func EncodeVarint(input []uint32) []byte {
	var result []byte
	for _, n := range input {
		number := []byte{byte(n % 128)}
		n /= 128
		for n > 0 {
			number = append(number, byte(n%128+128))
			n /= 128
		}

		for i, j := 0, len(number)-1; i < j; i, j = i+1, j-1 {
			number[i], number[j] = number[j], number[i]
		}

		result = append(result, number...)
	}
	return result
}

func DecodeVarint(input []byte) ([]uint32, error) {
	var numbers []uint32
	var currentBytes []byte

	for _, b := range input {
		currentBytes = append(currentBytes, b)

		if b < 128 {
			var number uint32
			len := len(currentBytes) - 1

			for i := 0; i <= len; i++ {
				n := currentBytes[i]
				if i != len {
					n -= 128
				}

				number += uint32(float64(n) * math.Pow(128, float64(len-i)))
			}

			numbers = append(numbers, number)
			currentBytes = []byte{}
		}
	}

	if len(currentBytes) > 0 {
		return nil, errors.New("sequence incomplete")
	}

	return numbers, nil
}
