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

	if len(l1) == len(l2) && reflect.DeepEqual(l1, l2) {
		return Equal
	}

	var reversed = false
	var a, b = l1, l2

	if len(a) < len(b) {
		a, b = b, a
		reversed = true
	}

	var result = Unequal
	for i := 0; i <= len(a)-len(b); i++ {
		if a[i] == b[0] && reflect.DeepEqual(a[i:len(b)+i], b) {

			result = SuperList
			break
		}
	}

	if reversed && result == SuperList {
		result = SubList
	}

	return result
}
