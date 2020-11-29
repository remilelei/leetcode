package Q0023_困难_合并K个升序链表

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

func TestQuestion23_1(t *testing.T) {
	listArray1 := []int{1, 4, 5}
	listArray2 := []int{1, 3, 4}
	listArray3 := []int{2, 6}
	listArrayAnswer := []int{1, 1, 2, 3, 4, 4, 5, 6}
	list1, _ := createList(listArray1, -1)
	list2, _ := createList(listArray2, -1)
	list3, _ := createList(listArray3, -1)
	inputLists := []*ListNode{list1, list2, list3}
	res := mergeKLists(inputLists)
	checkList(t, res, listArrayAnswer)
}

func TestQuestion23_2(t *testing.T) {
	// listArray1 := []int{1, 4, 5}
	// listArray2 := []int{1, 3, 4}
	// listArray3 := []int{2, 6}
	listArrayAnswer := []int{}
	// list1, _ := createList(listArray1, -1)
	// list2, _ := createList(listArray2, -1)
	// list3, _ := createList(listArray3, -1)
	inputLists := []*ListNode{}
	res := mergeKLists(inputLists)
	checkList(t, res, listArrayAnswer)
}
func TestQuestion23_3(t *testing.T) {
	listArray1 := []int{1}
	// listArray2 := []int{1, 3, 4}
	// listArray3 := []int{2, 6}
	listArrayAnswer := []int{1}
	list1, _ := createList(listArray1, -1)
	// list2, _ := createList(listArray2, -1)
	// list3, _ := createList(listArray3, -1)
	inputLists := []*ListNode{nil, list1}
	res := mergeKLists(inputLists)
	checkList(t, res, listArrayAnswer)
}
