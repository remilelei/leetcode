package Q0017_中等_电话号码的字母组合

import "testing"

func TestQuestion16_1(t *testing.T) {
	inputDigits := "23"
	// answer := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	res := letterCombinations(inputDigits)
	t.Logf("res=%v", res)
}
