package Q0009_简单_回文数

import "testing"

func TestQuestion9_1(t *testing.T) {
	inputS := 121
	answer := true
	res := isPalindrome(inputS)
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}

func TestQuestion9_2(t *testing.T) {
	inputS := -121
	answer := false
	res := isPalindrome(inputS)
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}

func TestQuestion9_3(t *testing.T) {
	inputS := 10
	answer := false
	res := isPalindrome(inputS)
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}
