package Q0010_困难_正则表达式匹配

import "testing"

func TestQuestion10_1(t *testing.T) {
	s := "aa"
	p := "a"
	answer := false
	res := isMatch(s, p)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion10_2(t *testing.T) {
	s := "aa"
	p := "a*"
	answer := true
	res := isMatch(s, p)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion10_3(t *testing.T) {
	s := "ab"
	p := ".*"
	answer := true
	res := isMatch(s, p)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion10_4(t *testing.T) {
	s := "aab"
	p := "c*a*b"
	answer := true
	res := isMatch(s, p)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion10_5(t *testing.T) {
	s := "mississippi"
	p := "mis*is*p*."
	answer := false
	res := isMatch(s, p)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion10_6(t *testing.T) {
	s := "aaa"
	p := "aaaa"
	answer := false
	res := isMatch(s, p)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
