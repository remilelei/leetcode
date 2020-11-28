package Q0020_简单_有效的括号

import "testing"

func TestQuestion20_1(t *testing.T) {
	inputS := "()"
	answer := true
	res := isValid(inputS)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion20_2(t *testing.T) {
	inputS := "()[]{}"
	answer := true
	res := isValid(inputS)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion20_3(t *testing.T) {
	inputS := "(]"
	answer := false
	res := isValid(inputS)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion20_4(t *testing.T) {
	inputS := "([)]"
	answer := false
	res := isValid(inputS)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion20_5(t *testing.T) {
	inputS := "{[]}"
	answer := true
	res := isValid(inputS)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion20_6(t *testing.T) {
	inputS := "{"
	answer := false
	res := isValid(inputS)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion20_7(t *testing.T) {
	inputS := "}"
	answer := false
	res := isValid(inputS)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
