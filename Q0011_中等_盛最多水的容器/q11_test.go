package Q0011_中等_盛最多水的容器

import "testing"

func TestQuestion11_1(t *testing.T) {
	inputHeight := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	answer := 49
	res := maxArea(inputHeight)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion11_2(t *testing.T) {
	inputHeight := []int{1, 1}
	answer := 1
	res := maxArea(inputHeight)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion11_3(t *testing.T) {
	inputHeight := []int{4, 3, 2, 1, 4}
	answer := 16
	res := maxArea(inputHeight)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion11_4(t *testing.T) {
	inputHeight := []int{1, 2, 1}
	answer := 2
	res := maxArea(inputHeight)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
