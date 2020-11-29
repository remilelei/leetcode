package Q0025_困难_K个一组翻转链表

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

	if len(data) == 0 && head != nil {
		t.Fatalf("wrong answer, answer is ampty, res=%v", head.Val)
	}

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

func TestQuestion25_1(t *testing.T) {
	listArray1 := []int{1, 2, 3, 4, 5}
	list1, _ := createList(listArray1, -1)
	inputK := 2
	res := reverseKGroup(list1, inputK)
	checkList(t, res, []int{2, 1, 4, 3, 5})
}
func TestQuestion25_2(t *testing.T) {
	listArray1 := []int{1, 2, 3, 4, 5}
	list1, _ := createList(listArray1, -1)
	inputK := 3
	res := reverseKGroup(list1, inputK)
	checkList(t, res, []int{3, 2, 1, 4, 5})
}
func TestQuestion25_3(t *testing.T) {
	listArray1 := []int{1, 2, 3, 4, 5}
	list1, _ := createList(listArray1, -1)
	inputK := 1
	res := reverseKGroup(list1, inputK)
	checkList(t, res, []int{1, 2, 3, 4, 5})
}
