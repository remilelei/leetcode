package Q0030_困难_串联所有单词的子串

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

func TestSolution_1(t *testing.T) {
	inputS := "barfoothefoobarman"
	inputWords := []string{"foo", "bar"}
	answer := []int{0, 9}
	res := findSubstring(inputS, inputWords)
	if !compareArray(res, answer) {
		t.Fatalf("answer wrong, res=%v, answer=%v", res, answer)
	}
}

func TestSolution_2(t *testing.T) {
	inputS := "wordgoodgoodgoodbestword"
	inputWords := []string{"word", "good", "best", "word"}
	answer := []int{}
	res := findSubstring(inputS, inputWords)
	if !compareArray(res, answer) {
		t.Fatalf("answer wrong, res=%v, answer=%v", res, answer)
	}
}
