package Q0027_简单_移除元素

import "testing"

func checkRes(arr1 []int, arr2 []int, length int) bool {
	for i := 0; i < length; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func TestSolution_1(t *testing.T) {
	inputNums := []int{3, 2, 2, 3}
	inputVal := 3
	res := removeElement(inputNums, inputVal)
	answer := []int{2, 2}
	answerLen := 2
	if res != answerLen {
		t.Fatalf("answer lenth is wrong, res=%v, but answerLen=%v", res, answerLen)
	}
	if !checkRes(inputNums, answer, res) {
		t.Fatalf("answer wrong, res=%v, answer=%v", inputNums, res)
	}
}
func TestSolution_2(t *testing.T) {
	inputNums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	inputVal := 2
	res := removeElement(inputNums, inputVal)
	answer := []int{0, 1, 3, 0, 4}
	answerLen := 5
	if res != answerLen {
		t.Fatalf("answer lenth is wrong, res=%v, but answerLen=%v", res, answerLen)
	}
	if !checkRes(inputNums, answer, res) {
		t.Fatalf("answer wrong, res=%v, answer=%v", inputNums, res)
	}
}
