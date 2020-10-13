package question0024

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。



示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var res *ListNode
	res = head
	if head == nil {
		return res
	}

	var nodeAfter *ListNode
	var nodeSwapA *ListNode
	var nodeSwapB *ListNode

	nodeSwapA = head
	nodeSwapB = nodeSwapA.Next

	for nodeSwapB != nil {
		nodeAfter = nodeSwapB.Next
		if nodeSwapA == head {
			res = nodeSwapB
		}

		// swap
		nodeSwapB.Next = nodeSwapA
		nodeSwapA.Next = nodeAfter

		if nodeAfter == nil || nodeAfter.Next == nil {
			break
		}

		// goto next pair
		nodeSwapA.Next = nodeAfter.Next
		nodeSwapA = nodeAfter
		nodeSwapB = nodeSwapA.Next
	}

	return res
}
