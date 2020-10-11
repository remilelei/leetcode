package question0142

import (
	"testing"
)

func createTestCase(t *testing.T, data []int, pos int) (head *ListNode, loop *ListNode) {
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

	cur.Next = loop
	return res, loop
}

func TestQuestion142_1(t *testing.T) {
	inputArr := []int{3, 2, 0, -4}
	inputPos := 1
	head, loop := createTestCase(t, inputArr, inputPos)

	res := detectCycle(head)
	if res != loop {
		resVal := -1
		if res != nil {
			resVal = res.Val
		}

		loopVal := -1
		if loop != nil {
			loopVal = loop.Val
		}

		t.Fatalf("res=%v loop=%v", resVal, loopVal)
	}
}
func TestQuestion142_2(t *testing.T) {
	inputArr := []int{1, 2}
	inputPos := 0
	head, loop := createTestCase(t, inputArr, inputPos)

	res := detectCycle(head)
	if res != loop {
		t.Fatalf("res=%v loop=%v", res, loop)
	}
}
func TestQuestion142_3(t *testing.T) {
	inputArr := []int{1}
	inputPos := -1
	head, loop := createTestCase(t, inputArr, inputPos)

	res := detectCycle(head)
	if res != loop {
		t.Fatalf("res=%v loop=%v", res, loop)
	}
}
