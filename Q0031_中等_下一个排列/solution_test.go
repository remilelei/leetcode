package Q0031_中等_下一个排列

import "testing"

func compareArray(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
func TestQuestion0031_1(t *testing.T) {
	input := []int{3, 2, 1}

	nextPermutation(input)

	answer := []int{1, 2, 3}

	if !compareArray(input, answer) {
		t.Fatalf("answer wrong, res=%v, answer=%v", input, answer)
	}
}
func TestQuestion0031_2(t *testing.T) {
	input := []int{1, 2, 3}

	nextPermutation(input)

	answer := []int{1, 3, 2}

	if !compareArray(input, answer) {
		t.Fatalf("answer wrong, res=%v, answer=%v", input, answer)
	}
}
func TestQuestion0031_3(t *testing.T) {
	input := []int{1, 1, 5}

	nextPermutation(input)

	answer := []int{1, 5, 1}

	if !compareArray(input, answer) {
		t.Fatalf("answer wrong, res=%v, answer=%v", input, answer)
	}
}
func TestQuestion0031_4(t *testing.T) {
	input := []int{7, 10, 9, 8, 7, 6}

	nextPermutation(input)

	answer := []int{8, 6, 7, 7, 9, 10}

	if !compareArray(input, answer) {
		t.Fatalf("answer wrong, res=%v, answer=%v", input, answer)
	}
}
