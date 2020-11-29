package Q0025_困难_K个一组翻转链表

/*
25. K 个一组翻转链表
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。



示例：

给你这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5



说明：

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(start *ListNode, boundary *ListNode) (newStart *ListNode, newEnd *ListNode) {
	reverseA := start
	if reverseA.Next == boundary {
		return reverseA, reverseA
	}
	reverseB := reverseA.Next
	reverseA.Next = nil
	newEnd = reverseA

	next := reverseB.Next

	for {
		reverseB.Next = reverseA
		if next == boundary || next == nil {
			newStart = reverseB
			return newStart, newEnd
		}

		reverseA = reverseB
		reverseB = next
		next = next.Next
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var newHead *ListNode
	reverseStart := head
	var lastReverseEnd *ListNode
	for reverseStart != nil {
		boundary := reverseStart
		for i := 0; i < k; i++ {
			if boundary != nil {
				boundary = boundary.Next
			} else {
				lastReverseEnd.Next = reverseStart
				return newHead
			}
		}
		newStart, newEnd := reverse(reverseStart, boundary)
		if lastReverseEnd == nil {
			newHead = newStart
		} else {
			lastReverseEnd.Next = newStart
		}
		lastReverseEnd = newEnd
		reverseStart = boundary
	}

	return newHead
}
