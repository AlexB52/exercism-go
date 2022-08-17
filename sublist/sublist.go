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
	if DeepEqual(l1, l2) {
		return Equal
	}

	var result Relation = Unequal
	var reversed bool

	if len(l1) < len(l2) {
		l1, l2 = l2, l1
		reversed = true
	}

	for i := 0; i <= len(l1)-len(l2); i++ {
		if DeepEqual(l1[i:len(l2)+i], l2) {
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
