package question0349两个数组的交集

import "testing"

func checkArray(res []int, answer []int) bool {
	if len(res) != len(answer) {
		return false
	}

	for i, v := range res {
		if answer[i] != v {
			return false
		}
	}
	return true
}

func TestQuestion349_1(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	answer := []int{2}
	res := intersection(nums1, nums2)
	if !checkArray(res, answer) {
		t.Fatalf("res=%v, but answer=%v", res, answer)
	}
}

func TestQuestion349_2(t *testing.T) {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	answer := []int{9, 4}
	res := intersection(nums1, nums2)
	if !checkArray(res, answer) {
		t.Fatalf("res=%v, but answer=%v", res, answer)
	}
}
