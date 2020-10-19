package question0844

import "testing"

func TestQuestion844_1(t *testing.T) {
	inputData1 := "ab#c"
	inputData2 := "ad#c"
	answer := true
	res := backspaceCompare(inputData1, inputData2)
	if answer != res {
		t.Fatalf("answer=%v, but res=%v", answer, res)
	}
}

func TestQuestion844_2(t *testing.T) {
	inputData1 := "ab##"
	inputData2 := "c#d#"
	answer := true
	res := backspaceCompare(inputData1, inputData2)
	if answer != res {
		t.Fatalf("answer=%v, but res=%v", answer, res)
	}
}

func TestQuestion844_3(t *testing.T) {
	inputData1 := "a##c"
	inputData2 := "#a#c"
	answer := true
	res := backspaceCompare(inputData1, inputData2)
	if answer != res {
		t.Fatalf("answer=%v, but res=%v", answer, res)
	}
}

func TestQuestion844_4(t *testing.T) {
	inputData1 := "a#c"
	inputData2 := "b"
	answer := false
	res := backspaceCompare(inputData1, inputData2)
	if answer != res {
		t.Fatalf("answer=%v, but res=%v", answer, res)
	}
}

func TestQuestion844_5(t *testing.T) {
	inputData1 := "bxj##tw"
	inputData2 := "bxj###tw"
	answer := false
	res := backspaceCompare(inputData1, inputData2)
	if answer != res {
		t.Fatalf("answer=%v, but res=%v", answer, res)
	}
}
