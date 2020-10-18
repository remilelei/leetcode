package question0019

import (
	"testing"
)

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

func TestQuestion19_1(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	deletePos := 2
	listNode, _ := createList(data, -1)
	res := removeNthFromEnd(listNode, deletePos)
	answer := []int{1, 2, 3, 5}
	checkList(t, res, answer)
}

func TestQuestion19_2(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	deletePos := 5
	listNode, _ := createList(data, -1)
	res := removeNthFromEnd(listNode, deletePos)
	answer := []int{2, 3, 4, 5}
	checkList(t, res, answer)
}
