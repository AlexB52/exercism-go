package railfence

import (
	"fmt"
	"testing"
)

func testCases(op func(string, int) string, cases []testCase, t *testing.T) {
	for _, tc := range cases {
		if actual := op(tc.message, tc.rails); actual != tc.expected {
			t.Fatalf("FAIL: %s\nExpected: %q\nActual: %q", tc.description, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func TestEncode(t *testing.T) { testCases(Encode, encodeTests, t) }
func TestDecode(t *testing.T) { testCases(Decode, decodeTests, t) }

func TestDiagonaleProjection(t *testing.T) {
	f2Tests := []struct {
		position int
		expected Node
	}{
		{0, Node{x: 0, y: 0}},
		{1, Node{x: 1, y: 1}},
		{2, Node{x: 2, y: 0}},
		{3, Node{x: 3, y: 1}},
		{4, Node{x: 4, y: 0}},
		{5, Node{x: 5, y: 1}},
		{6, Node{x: 6, y: 0}},
		{7, Node{x: 7, y: 1}},
		{8, Node{x: 8, y: 0}},
		{9, Node{x: 9, y: 1}},
		{10, Node{x: 10, y: 0}},
		{11, Node{x: 11, y: 1}},
	}

	f2 := DiagonaleProjection(2)
	for _, test := range f2Tests {
		t.Run(fmt.Sprintf("f2 with position: %d", test.position), func(t *testing.T) {
			actual := f2(test.position)
			if actual != test.expected {
				t.Fatalf("FAIL - Expected: %v Actual: %v", test.expected, actual)
			}
		})
	}

	f3Tests := []struct {
		position int
		expected Node
	}{
		{0, Node{x: 0, y: 0}},
		{1, Node{x: 1, y: 1}},
		{2, Node{x: 2, y: 2}},
		{3, Node{x: 3, y: 1}},
		{4, Node{x: 4, y: 0}},
		{5, Node{x: 5, y: 1}},
		{6, Node{x: 6, y: 2}},
		{7, Node{x: 7, y: 1}},
		{8, Node{x: 8, y: 0}},
		{9, Node{x: 9, y: 1}},
		{10, Node{x: 10, y: 2}},
		{11, Node{x: 11, y: 1}},
	}

	f3 := DiagonaleProjection(3)
	for _, test := range f3Tests {
		t.Run(fmt.Sprintf("f3 with position: %d", test.position), func(t *testing.T) {
			actual := f3(test.position)
			if actual != test.expected {
				t.Fatalf("FAIL - Expected: %v Actual: %v", test.expected, actual)
			}
		})
	}

	f4Tests := []struct {
		position int
		expected Node
	}{
		{0, Node{x: 0, y: 0}},
		{1, Node{x: 1, y: 1}},
		{2, Node{x: 2, y: 2}},
		{3, Node{x: 3, y: 3}},
		{4, Node{x: 4, y: 2}},
		{5, Node{x: 5, y: 1}},
		{6, Node{x: 6, y: 0}},
		{7, Node{x: 7, y: 1}},
		{8, Node{x: 8, y: 2}},
		{9, Node{x: 9, y: 3}},
		{10, Node{x: 10, y: 2}},
		{11, Node{x: 11, y: 1}},
	}

	f4 := DiagonaleProjection(4)
	for _, test := range f4Tests {
		t.Run(fmt.Sprintf("f4 with position: %d", test.position), func(t *testing.T) {
			actual := f4(test.position)
			if actual != test.expected {
				t.Fatalf("FAIL - Expected: %v Actual: %v", test.expected, actual)
			}
		})
	}

	f5Tests := []struct {
		position int
		expected Node
	}{
		{0, Node{x: 0, y: 0}},
		{1, Node{x: 1, y: 1}},
		{2, Node{x: 2, y: 2}},
		{3, Node{x: 3, y: 3}},
		{4, Node{x: 4, y: 4}},
		{5, Node{x: 5, y: 3}},
		{6, Node{x: 6, y: 2}},
		{7, Node{x: 7, y: 1}},
		{8, Node{x: 8, y: 0}},
		{9, Node{x: 9, y: 1}},
		{10, Node{x: 10, y: 2}},
		{11, Node{x: 11, y: 3}},
	}

	f5 := DiagonaleProjection(5)
	for _, test := range f5Tests {
		t.Run(fmt.Sprintf("f5 with position: %d", test.position), func(t *testing.T) {
			actual := f5(test.position)
			if actual != test.expected {
				t.Fatalf("FAIL - Expected: %v Actual: %v", test.expected, actual)
			}
		})
	}
}
