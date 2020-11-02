package question0763划分字母区间

import "testing"

func checkArray(res []int, answer []int) bool {
	if len(res) != len(answer) {
		return false
	}

	for i, v := range res {
		if answer[i] != v {
			return false
		}
	}
	return true
}

func TestQuestion763_1(t *testing.T) {
	input := "ababcbacadefegdehijhklij"
	answer := []int{9, 7, 8}
	res := partitionLabels(input)
	if !checkArray(res, answer) {
		t.Fatalf("res=%v, but answer=%v", res, answer)
	}
}

func TestQuestion763_2(t *testing.T) {
	input := "caedbdedda"
	answer := []int{1, 9}
	res := partitionLabels(input)
	if !checkArray(res, answer) {
		t.Fatalf("res=%v, but answer=%v", res, answer)
	}
}
