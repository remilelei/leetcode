package question0001

import "testing"

func TestQuestion1_1(t *testing.T) {
	inputArr := []int{2, 7, 11, 15}
	inputTarget := 9
	result := Q1TwoSum(inputArr, inputTarget)
	answerResult := [2]int{0, 1}

	if len(answerResult) != len(result) {
		t.Fail()
	}

	for index, value := range answerResult {
		if value != result[index] {
			t.Fatalf("answer=%v, result=%v", answerResult, result)
		}
	}
}
func TestQuestion1_2(t *testing.T) {
	inputArr := []int{2, 7, 11, 15}
	inputTarget := 13
	result := Q1TwoSum(inputArr, inputTarget)
	answerResult := [2]int{0, 2}

	if len(answerResult) != len(result) {
		t.Fail()
	}

	for index, value := range answerResult {
		if value != result[index] {
			t.Fatalf("answer=%v, result=%v", answerResult, result)
		}
	}
}
func TestQuestion1_3(t *testing.T) {
	inputArr := []int{2, 7, 11, 15}
	inputTarget := 18
	result := Q1TwoSum(inputArr, inputTarget)
	answerResult := [2]int{1, 2}

	if len(answerResult) != len(result) {
		t.Fail()
	}

	for index, value := range answerResult {
		if value != result[index] {
			t.Fatalf("answer=%v, result=%v", answerResult, result)
		}
	}
}
