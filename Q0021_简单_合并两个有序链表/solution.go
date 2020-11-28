package Q0021_简单_合并两个有序链表

/*
21. 合并两个有序链表
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。



示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var ret *ListNode
	var retCur *ListNode
	for l1 != nil || l2 != nil {
		var minCur *ListNode
		if l1 == nil || (l2 != nil && l2.Val < l1.Val) {
			minCur = l2
			l2 = l2.Next
		} else {
			minCur = l1
			l1 = l1.Next
		}
		if ret == nil {
			ret = minCur
			retCur = ret
		} else {
			retCur.Next = minCur
			retCur = retCur.Next
		}
	}

	if retCur != nil {
		retCur.Next = nil
	}

	return ret
}
