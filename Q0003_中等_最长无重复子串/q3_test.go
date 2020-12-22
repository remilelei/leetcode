package Q0003_中等_最长无重复子串

import "testing"

func TestQuestion3_1(t *testing.T) {
	data := "bbbbb"
	res := lengthOfLongestSubstring(data)
	answer := 1
	if res != answer {
		t.Fatalf("[%v] answer=%v but res=%v", data, answer, res)
	}
}

func TestQuestion3_2(t *testing.T) {
	data := "abcabcbb"
	res := lengthOfLongestSubstring(data)
	answer := 3
	if res != answer {
		t.Fatalf("[%v] answer=%v but res=%v", data, answer, res)
	}
}

func TestQuestion3_3(t *testing.T) {
	data := "pwwkew"
	res := lengthOfLongestSubstring(data)
	answer := 3
	if res != answer {
		t.Fatalf("[%v] answer=%v but res=%v", data, answer, res)
	}
}

func TestQuestion3_4(t *testing.T) {
	data := "aab"
	res := lengthOfLongestSubstring(data)
	answer := 2
	if res != answer {
		t.Fatalf("[%v] answer=%v but res=%v", data, answer, res)
	}
}

func TestQuestion3_5(t *testing.T) {
	data := "dvdf"
	res := lengthOfLongestSubstring(data)
	answer := 3
	if res != answer {
		t.Fatalf("[%v] answer=%v but res=%v", data, answer, res)
	}
}

func TestQuestion3_6(t *testing.T) {
	data := "dvdfabcdad"
	res := lengthOfLongestSubstring(data)
	answer := 6
	if res != answer {
		t.Fatalf("[%v] answer=%v but res=%v", data, answer, res)
	}
}

func TestQuestion3_7(t *testing.T) {
	data := "abba"
	res := lengthOfLongestSubstring(data)
	answer := 2
	if res != answer {
		t.Fatalf("[%v] answer=%v but res=%v", data, answer, res)
	}
}
