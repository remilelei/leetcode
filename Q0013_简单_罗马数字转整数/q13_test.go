package Q0013_简单_罗马数字转整数

import "testing"

func TestQuestion13_1(t *testing.T) {
	inputS := "III"
	answer := 3
	res := romanToInt(inputS)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion13_2(t *testing.T) {
	inputS := "IV"
	answer := 4
	res := romanToInt(inputS)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion13_3(t *testing.T) {
	inputS := "IX"
	answer := 9
	res := romanToInt(inputS)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion13_4(t *testing.T) {
	inputS := "LVIII"
	answer := 58
	res := romanToInt(inputS)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion13_5(t *testing.T) {
	inputS := "MCMXCIV"
	answer := 1994
	res := romanToInt(inputS)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
