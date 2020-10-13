package question0024

import "testing"

func createTestCase(data []int, pos int) (head *ListNode, loop *ListNode) {
	var res *ListNode
	var cur *ListNode
	for index, val := range data {
		var node ListNode
		node.Val = val

		if res == nil {
			res = &node
		}

		if index == pos {
			loop = &node
		}

		if cur == nil {
			cur = &node
		} else {
			cur.Next = &node
			cur = &node
		}
	}

	if cur != nil {
		cur.Next = loop
	}
	return res, loop
}

func checkList(t *testing.T, head *ListNode, data []int) {
	cur := head

	for index, val := range data {
		if cur == nil {
			t.Fatalf("list go to the end, index=%v", index)
		}
		if cur.Val != val {
			t.Fatalf("wrong answer, cur.Val=%v, data[%v]=%v", cur.Val, index, val)
		} else {
			cur = cur.Next
		}
	}
}

func TestQuestion24_1(t *testing.T) {
	data := []int{1, 2, 3, 4}
	list, _ := createTestCase(data, -1)
	newList := swapPairs(list)
	answer := []int{2, 1, 4, 3}
	checkList(t, newList, answer)
}
