package Q0012_中等_整数转罗马数字

import "testing"

func TestQuestion12_1(t *testing.T) {
	inputNum := 3
	answer := "III"
	res := intToRoman(inputNum)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion12_2(t *testing.T) {
	inputNum := 4
	answer := "IV"
	res := intToRoman(inputNum)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion12_3(t *testing.T) {
	inputNum := 9
	answer := "IX"
	res := intToRoman(inputNum)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion12_4(t *testing.T) {
	inputNum := 58
	answer := "LVIII"
	res := intToRoman(inputNum)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion12_5(t *testing.T) {
	inputNum := 1994
	answer := "MCMXCIV"
	res := intToRoman(inputNum)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQuestion12_6(t *testing.T) {
	inputNum := 5000
	answer := "FUCK"
	res := intToRoman(inputNum)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
