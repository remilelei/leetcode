package question0530

import (
	"math"
	"sort"
)

/**
给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。

示例：

输入：

   1
    \
     3
    /
   2

输出：
1

解释：
最小绝对差为 1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。

提示：

树中至少有 2 个节点。
本题与 783 https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/ 相同

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// TreeNode 题目中二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func transTree2Slice(node *TreeNode, slice *[]int) {
	if node != nil {
		(*slice) = append(*slice, node.Val)
		transTree2Slice(node.Left, slice)
		transTree2Slice(node.Right, slice)
	}
}

func getMinimumDifference(root *TreeNode) int {
	slice := []int{}
	transTree2Slice(root, &slice)
	sort.Sort(sort.IntSlice(slice))
	res := -1
	for i := len(slice) - 2; i >= 0; i-- {
		tempRes := slice[i+1] - slice[i]
		if res == -1 || tempRes < res {
			res = tempRes
		}
	}
	return res
}

/**
本题是一个【二叉搜索树】，二叉搜索树以中序遍历能得到一个升序数组，不用再次排序
在这个升序数组中可以找到最小的差
*/
func getMinimumDifferenceOffical(root *TreeNode) int {
	answer, pre := math.MaxInt32, -1
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 && node.Val-pre < answer {
			answer = node.Val - pre
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return answer
}
