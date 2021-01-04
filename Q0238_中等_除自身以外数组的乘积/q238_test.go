package Q0238_中等_除自身以外数组的乘积

import "testing"

func TestQuestion238_1(t *testing.T) {
	inputNums := []int{1, 2, 3, 4}
	res := productExceptSelf(inputNums)
	answer := []int{24, 12, 8, 6}

	if len(res) != len(answer) {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}

	for i, v := range res {
		if v != answer[i] {
			t.Fatalf("answer=%v but res=%v", answer, res)
		}
	}
}
