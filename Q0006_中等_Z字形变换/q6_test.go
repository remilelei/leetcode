package Q0006_中等_Z字形变换

import "testing"

func TestQ6_1(t *testing.T) {
	inputS := "LEETCODEISHIRING"
	inputNumRows := 3
	res := convert(inputS, inputNumRows)
	answer := "LCIRETOESIIGEDHN"
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}

func TestQ6_2(t *testing.T) {
	inputS := "LEETCODEISHIRING"
	inputNumRows := 4
	res := convert(inputS, inputNumRows)
	answer := "LDREOEIIECIHNTSG"
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}

func TestQ6_3(t *testing.T) {
	inputS := "A"
	inputNumRows := 1
	res := convert(inputS, inputNumRows)
	answer := "A"
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}

func TestQ6_4(t *testing.T) {
	inputS := "0123456789"
	inputNumRows := 2
	res := convert(inputS, inputNumRows)
	answer := "0246813579"
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}
