package twobucket

import (
	"errors"
)

func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (string, int, int, error) {
	err := ValidateInputs(sizeBucketOne, sizeBucketTwo, goalAmount, startBucket)
	if err != nil {
		return "", 0, 0, err
	}

	b1, b2 := &Bucket{name: "one", size: sizeBucketOne}, &Bucket{name: "two", size: sizeBucketTwo}
	if startBucket != "one" {
		b1, b2 = b2, b1
	}

	var rounds int
	for b1.quantity != goalAmount && b2.quantity != goalAmount {
		rounds++
		switch {
		case b1.IsEmpty():
			b1.Fill()
		case b2.IsSize(goalAmount):
			b2.Fill()
		case b2.IsFull():
			b2.Empty()
		default:
			b1.PourInto(b2)
		}
	}

	goalBucket, otherBucket := FindBuckets(b1, b2, goalAmount)

	return goalBucket.name, rounds, otherBucket.quantity, nil
}

func ValidateInputs(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) error {
	if sizeBucketOne == 0 || sizeBucketTwo == 0 {
		return errors.New("invalid bucket size")
	}

	if goalAmount == 0 {
		return errors.New("invalid amount")
	}

	if startBucket != "one" && startBucket != "two" {
		return errors.New("invalid start bucket")
	}

	if sizeBucketOne > 1 && sizeBucketTwo%sizeBucketOne == 0 {
		return errors.New("impossible solution")
	}

	return nil
}

func FindBuckets(b1, b2 *Bucket, goal int) (*Bucket, *Bucket) {
	if b1.quantity == goal {
		return b1, b2
	} else {
		return b2, b1
	}
}
