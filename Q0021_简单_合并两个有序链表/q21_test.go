package Q0021_简单_合并两个有序链表

import "testing"

func createList(data []int, pos int) (head *ListNode, loop *ListNode) {
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

func TestQuestion21_1(t *testing.T) {
	inputS := "()"
	answer := true
	res := mergeTwoLists(inputS)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
