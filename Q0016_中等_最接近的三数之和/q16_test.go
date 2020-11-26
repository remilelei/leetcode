package Q0016_中等_最接近的三数之和

import "testing"

func TestQuestion16_1(t *testing.T) {
	inputNums := []int{-1, 2, 1, -4}
	inputTarget := 1
	answer := 2
	res := threeSumClosest(inputNums, inputTarget)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
