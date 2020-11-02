package Q0004_困难_寻找两个正序数组的中位数

import "testing"

func TestQ4_1(t *testing.T) {
	inputData1 := []int{1, 3}
	inputData2 := []int{2}
	answer := 2.0
	res := findMedianSortedArrays(inputData1, inputData2)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQ4_2(t *testing.T) {
	inputData1 := []int{1, 2}
	inputData2 := []int{3, 4}
	answer := 2.5
	res := findMedianSortedArrays(inputData1, inputData2)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQ4_3(t *testing.T) {
	inputData1 := []int{0, 0}
	inputData2 := []int{0, 0}
	answer := 0.0
	res := findMedianSortedArrays(inputData1, inputData2)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQ4_4(t *testing.T) {
	inputData1 := []int{}
	inputData2 := []int{1}
	answer := 1.0
	res := findMedianSortedArrays(inputData1, inputData2)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQ4_5(t *testing.T) {
	inputData1 := []int{2}
	inputData2 := []int{}
	answer := 2.0
	res := findMedianSortedArrays(inputData1, inputData2)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQ4_6(t *testing.T) {
	inputData1 := []int{2}
	inputData2 := []int{1, 3, 4}
	answer := 2.5
	res := findMedianSortedArrays(inputData1, inputData2)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQ4_7(t *testing.T) {
	inputData1 := []int{1}
	inputData2 := []int{2, 3, 4}
	answer := 2.5
	res := findMedianSortedArrays(inputData1, inputData2)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQ4_8(t *testing.T) {
	inputData1 := []int{3}
	inputData2 := []int{-2, -1}
	answer := -1.0
	res := findMedianSortedArrays(inputData1, inputData2)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQ4_9(t *testing.T) {
	inputData1 := []int{0, 0, 0, 0, 0}
	inputData2 := []int{-1, 0, 0, 0, 0, 0, 1}
	answer := 0.0
	res := findMedianSortedArrays(inputData1, inputData2)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQ4_10(t *testing.T) {
	inputData1 := []int{1, 2}
	inputData2 := []int{-1, 3}
	answer := 1.5
	res := findMedianSortedArrays(inputData1, inputData2)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQ4_11(t *testing.T) {
	inputData1 := []int{1, 2}
	inputData2 := []int{1, 1}
	answer := 1.0
	res := findMedianSortedArrays(inputData1, inputData2)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQ4_12(t *testing.T) {
	inputData1 := []int{2}
	inputData2 := []int{1, 3, 4, 5}
	answer := 3.0
	res := findMedianSortedArrays(inputData1, inputData2)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
