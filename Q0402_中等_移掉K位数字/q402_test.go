package Q0402_中等_移掉K位数字

import "testing"

func TestQuestion402_1(t *testing.T) {
	inputNum := "1432219"
	inputK := 3
	res := removeKdigits(inputNum, inputK)
	answer := "1219"
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQuestion402_2(t *testing.T) {
	inputNum := "10200"
	inputK := 1
	res := removeKdigits(inputNum, inputK)
	answer := "200"
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQuestion402_3(t *testing.T) {
	inputNum := "10"
	inputK := 2
	res := removeKdigits(inputNum, inputK)
	answer := "0"
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQuestion402_4(t *testing.T) {
	inputNum := "10200"
	inputK := 2
	res := removeKdigits(inputNum, inputK)
	answer := "0"
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQuestion402_5(t *testing.T) {
	inputNum := "112"
	inputK := 1
	res := removeKdigits(inputNum, inputK)
	answer := "11"
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQuestion402_6(t *testing.T) {
	inputNum := "10"
	inputK := 1
	res := removeKdigits(inputNum, inputK)
	answer := "0"
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQuestion402_7(t *testing.T) {
	inputNum := "1173"
	inputK := 2
	res := removeKdigits(inputNum, inputK)
	answer := "11"
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQuestion402_8(t *testing.T) {
	inputNum := "12345"
	inputK := 0
	res := removeKdigits(inputNum, inputK)
	answer := "12345"
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQuestion402_9(t *testing.T) {
	inputNum := "1234567890"
	inputK := 9
	res := removeKdigits(inputNum, inputK)
	answer := "0"
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
