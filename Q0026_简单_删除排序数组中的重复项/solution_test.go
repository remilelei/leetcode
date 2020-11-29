package Q0026_简单_删除排序数组中的重复项

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
	inputNums := []int{1, 1, 2}
	answer := []int{1, 2}
	res := removeDuplicates(inputNums)
	answerLen := 2
	if res != answerLen {
		t.Fatalf("answer lenth is wrong, res=%v, but answerLen=%v", res, answerLen)
	}
	if !checkRes(inputNums, answer, res) {
		t.Fatalf("answer wrong, res=%v, answer=%v", inputNums, res)
	}
}
func TestSolution_2(t *testing.T) {
	inputNums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	answer := []int{0, 1, 2, 3, 4}
	res := removeDuplicates(inputNums)
	answerLen := 5
	if res != answerLen {
		t.Fatalf("answer lenth is wrong, res=%v, but answerLen=%v", res, answerLen)
	}
	if !checkRes(inputNums, answer, res) {
		t.Fatalf("answer wrong, res=%v, answer=%v", inputNums, res)
	}
}
