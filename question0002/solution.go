package question0002

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	curRes := res
	curL1 := l1
	curL2 := l2
	preAdd := false

	for curL1 != nil || curL2 != nil {
		bitSum := 0

		if curL1 != nil {
			bitSum += curL1.Val
			curL1 = curL1.Next
		}

		if curL2 != nil {
			bitSum += curL2.Val
			curL2 = curL2.Next
		}

		if preAdd {
			bitSum += 1
			preAdd = false
		}

		if bitSum >= 10 {
			preAdd = true
			bitSum -= 10
		}

		var node ListNode
		node.Val = bitSum

		if res == nil || curRes == nil {
			res = &node
			curRes = &node
		} else {
			curRes.Next = &node
			curRes = curRes.Next
		}
	}

	if preAdd {
		if curRes != nil {
			curRes.Next = &ListNode{Val: 1}
		}
	}

	return res
}
