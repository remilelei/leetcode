package question0142

/**
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。

说明：不允许修改给定的链表。

进阶：

你是否可以不用额外空间解决此题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/linked-list-cycle-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// ListNode 题目中的节点类型
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
从起点放出速度为1、速度为2的两探针，若相遇则必然有环
1. 环外长度为a
2. 环内·环起点到探针相遇点，长度为b
3. 环内·探针相遇点到环起点，长度为c

 2*(a+b+m*(b+c))=a+b+n*(b+c)
 a+b=(n-2m)(b+c)
 n-2m简化为任一未知数x
 a+b=x*(b+c)=(x-1)*(b+c)+c
 因(b+c)为一个环，相当于未移动，则有 a=c
 所以分别从探针相遇点与起点再次放出等速两探针，相遇点则为环起点
*/
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			// 发生异速相遇，能相遇则必有环
			calNode1 := head
			calNode2 := fast
			for calNode1 != calNode2 {
				calNode1 = calNode1.Next
				calNode2 = calNode2.Next
			}
			return calNode1
		}
	}

	return nil
}
