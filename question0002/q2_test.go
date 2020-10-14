package question0002

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

func TestQuestion2_1(t *testing.T) {
	inputData1 := []int{2, 4, 3}
	inputData2 := []int{5, 6, 4}
	tempData1, _ := createTestCase(inputData1, -1)
	tempData2, _ := createTestCase(inputData2, -1)
	res := addTwoNumbers(tempData1, tempData2)
	answer := []int{7, 0, 8}
	checkList(t, res, answer)
}

func TestQuestion2_2(t *testing.T) {
	inputData1 := []int{2, 4, 5}
	inputData2 := []int{5, 6, 4}
	tempData1, _ := createTestCase(inputData1, -1)
	tempData2, _ := createTestCase(inputData2, -1)
	res := addTwoNumbers(tempData1, tempData2)
	answer := []int{7, 0, 0, 1}
	checkList(t, res, answer)
}

func TestQuestion2_3(t *testing.T) {
	inputData1 := []int{2, 4, 5}
	inputData2 := []int{0}
	tempData1, _ := createTestCase(inputData1, -1)
	tempData2, _ := createTestCase(inputData2, -1)
	res := addTwoNumbers(tempData1, tempData2)
	answer := []int{2, 4, 5}
	checkList(t, res, answer)
}
