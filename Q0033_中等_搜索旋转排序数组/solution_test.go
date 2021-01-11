package Q0033_中等_搜索旋转排序数组

import "testing"

func TestQuestion0033_1(t *testing.T) {
	input := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0

	res := search(input, target)

	answer := 4

	if res != answer {
		t.Fatalf("answer wrong, res=%v, answer=%v", input, answer)
	}
}
func TestQuestion0033_2(t *testing.T) {
	input := []int{4, 5, 6, 7, 0, 1, 2}
	target := 3

	res := search(input, target)

	answer := -1

	if res != answer {
		t.Fatalf("answer wrong, res=%v, answer=%v", input, answer)
	}
}
func TestQuestion0033_3(t *testing.T) {
	input := []int{1}
	target := 0

	res := search(input, target)

	answer := -1

	if res != answer {
		t.Fatalf("answer wrong, res=%v, answer=%v", input, answer)
	}
}
func TestQuestion0033_4(t *testing.T) {
	input := []int{6, 6, 6, 6}
	target := 7

	res := search(input, target)

	answer := -1

	if res != answer {
		t.Fatalf("answer wrong, res=%v, answer=%v", input, answer)
	}
}
func TestQuestion0033_5(t *testing.T) {
	input := []int{1, 3}
	target := 0

	res := search(input, target)

	answer := -1

	if res != answer {
		t.Fatalf("answer wrong, res=%v, answer=%v", input, answer)
	}
}
