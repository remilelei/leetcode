package Q0023_困难_合并K个升序链表

/*
23. 合并K个升序链表
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。



示例 1：

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
示例 2：

输入：lists = []
输出：[]
示例 3：

输入：lists = [[]]
输出：[]


提示：

k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] 按 升序 排列
lists[i].length 的总和不超过 10^4

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeList(listA *ListNode, listB *ListNode) *ListNode {
	var ret *ListNode
	var retCur *ListNode
	for listA != nil && listB != nil {
		var minCur *ListNode
		if listA.Val < listB.Val {
			minCur = listA
			listA = listA.Next
		} else {
			minCur = listB
			listB = listB.Next
		}

		if ret == nil {
			ret = minCur
			retCur = ret
		} else {
			retCur.Next = minCur
			retCur = retCur.Next
		}
	}

	if listA != nil {
		if ret == nil {
			ret = listA
		} else {
			retCur.Next = listA
		}
	}
	if listB != nil {
		if ret == nil {
			ret = listB
		} else {
			retCur.Next = listB
		}
	}

	return ret
}

func mergeLists(lists []*ListNode, start int, end int) *ListNode {
	if start < 0 || end >= len(lists) {
		return nil
	}

	if start == end {
		return lists[start]
	}

	if end == start+1 {
		return mergeList(lists[start], lists[end])
	}

	mid := (start + end) / 2
	ret1 := mergeLists(lists, start, mid-1)
	ret2 := mergeLists(lists, mid, end)
	return mergeList(ret1, ret2)
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	return mergeLists(lists, 0, len(lists)-1)
}
