package sublist

import (
	"reflect"
)

type Relation string

const (
	SubList   Relation = "sublist"
	SuperList Relation = "superlist"
	Equal     Relation = "equal"
	Unequal   Relation = "unequal"
)

func Sublist(l1, l2 []int) Relation {
	if len(l1) != 0 && len(l2) == 0 {
		return SuperList
	}

	if len(l1) == 0 && len(l2) != 0 {
		return SubList
	}

	if len(l1) == len(l2) && DeepEqual(l1, l2) {
		return Equal
	}

	var result Relation = Unequal
	var reversed bool
	var a, b = l1, l2

	if len(a) < len(b) {
		a, b = b, a
		reversed = true
	}

	for i := 0; i <= len(a)-len(b); i++ {
		if a[i] == b[0] && DeepEqual(a[i:len(b)+i], b) {
			result = SuperList
			if reversed {
				result = SubList
			}
			break
		}
	}

	return result
}

func DeepEqual(a, b []int) bool {
	return reflect.DeepEqual(a, b)
}
