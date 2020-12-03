package Q0029_中等_两数相除

import "testing"

func TestSolution_1(t *testing.T) {
	inputDividend := 10
	inputDivisor := 3
	answer := 3
	res := divide(inputDividend, inputDivisor)
	if res != answer {
		t.Fatalf("answer wrong, res=%v, answer=%v", res, answer)
	}
}

func TestSolution_2(t *testing.T) {
	inputDividend := 7
	inputDivisor := -3
	answer := -2
	res := divide(inputDividend, inputDivisor)
	if res != answer {
		t.Fatalf("answer wrong, res=%v, answer=%v", res, answer)
	}
}

func TestSolution_3(t *testing.T) {
	inputDividend := 7
	inputDivisor := -3
	answer := -2
	res := divide(inputDividend, inputDivisor)
	if res != answer {
		t.Fatalf("answer wrong, res=%v, answer=%v", res, answer)
	}
}

func TestSolution_4(t *testing.T) {
	inputDividend := 1
	inputDivisor := 1
	answer := 1
	res := divide(inputDividend, inputDivisor)
	if res != answer {
		t.Fatalf("answer wrong, res=%v, answer=%v", res, answer)
	}
}

func TestSolution_5(t *testing.T) {
	inputDividend := 1
	inputDivisor := 2
	answer := 0
	res := divide(inputDividend, inputDivisor)
	if res != answer {
		t.Fatalf("answer wrong, res=%v, answer=%v", res, answer)
	}
}

func TestSolution_6(t *testing.T) {
	inputDividend := -2147483648
	inputDivisor := -1
	answer := 2147483647
	res := divide(inputDividend, inputDivisor)
	if res != answer {
		t.Fatalf("answer wrong, res=%v, answer=%v", res, answer)
	}
}
