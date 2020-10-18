package question0019

/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := head

	fast := head

	slow := head // 删他
	beforeSlow := head
	afterSlow := head

	slowMoveable := false
	beforeSlowMoveable := false
	afterSlowMoveable := false

	curPos := 0
	for fast != nil {
		if curPos == (n - 1) {
			afterSlowMoveable = true
		}
		if curPos == n {
			slowMoveable = true
		}
		if curPos == (n + 1) {
			beforeSlowMoveable = true
		}

		fast = fast.Next
		if slowMoveable && slow != nil {
			slow = slow.Next
		}
		if beforeSlowMoveable && beforeSlow != nil {
			beforeSlow = beforeSlow.Next
		}
		if afterSlowMoveable && afterSlow != nil {
			afterSlow = afterSlow.Next
		}

		curPos++
	}

	if slow == head {
		res = head.Next
	} else {
		beforeSlow.Next = afterSlow
	}

	return res
}
