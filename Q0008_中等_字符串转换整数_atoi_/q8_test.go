package Q0008_中等_字符串转换整数_atoi_

import "testing"

func TestQuestion8_1(t *testing.T) {
	inputS := "42"
	answer := 42
	res := myAtoi(inputS)
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}

func TestQuestion8_2(t *testing.T) {
	inputS := "   -42"
	answer := -42
	res := myAtoi(inputS)
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}

func TestQuestion8_3(t *testing.T) {
	inputS := "4193 with words"
	answer := 4193
	res := myAtoi(inputS)
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}

func TestQuestion8_4(t *testing.T) {
	inputS := "words and 987"
	answer := 0
	res := myAtoi(inputS)
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}

func TestQuestion8_5(t *testing.T) {
	inputS := "-91283472332"
	answer := -2147483648
	res := myAtoi(inputS)
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}

func TestQuestion8_6(t *testing.T) {
	inputS := "9223372036854775808"
	answer := 2147483647
	res := myAtoi(inputS)
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}
