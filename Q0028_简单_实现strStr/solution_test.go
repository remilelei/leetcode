package Q0028_简单_实现strStr

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
	inputHaystack := "hello"
	inputNeedle := "ll"
	res := strStr(inputHaystack, inputNeedle)
	answer := 2
	if res != answer {
		t.Fatalf("answer wrong, res=%v, answer=%v", res, res)
	}
}
func TestSolution_2(t *testing.T) {
	inputHaystack := "aaaaa"
	inputNeedle := "bba"
	res := strStr(inputHaystack, inputNeedle)
	answer := -1
	if res != answer {
		t.Fatalf("answer wrong, res=%v, answer=%v", res, res)
	}
}
func TestSolution_3(t *testing.T) {
	inputHaystack := "a"
	inputNeedle := ""
	res := strStr(inputHaystack, inputNeedle)
	answer := 0
	if res != answer {
		t.Fatalf("answer wrong, res=%v, answer=%v", res, res)
	}
}
func TestSolution_4(t *testing.T) {
	inputHaystack := "abc"
	inputNeedle := "c"
	res := strStr(inputHaystack, inputNeedle)
	answer := 2
	if res != answer {
		t.Fatalf("answer wrong, res=%v, answer=%v", res, res)
	}
}
