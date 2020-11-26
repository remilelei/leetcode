package Q0014_简单_最长公共前缀

import "testing"

func TestQuestion14_1(t *testing.T) {
	inputStrs := []string{"flower", "flow", "flight"}
	answer := "fl"
	res := longestCommonPrefix(inputStrs)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion14_2(t *testing.T) {
	inputStrs := []string{"dog", "racecar", "car"}
	answer := ""
	res := longestCommonPrefix(inputStrs)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
