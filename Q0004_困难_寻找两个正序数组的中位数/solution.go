package Q0004_困难_寻找两个正序数组的中位数

import "math"

/*
给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。

进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？



示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
示例 3：

输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000
示例 4：

输入：nums1 = [], nums2 = [1]
输出：1.00000
示例 5：

输入：nums1 = [2], nums2 = []
输出：2.00000


提示：

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	K := totalLen / 2 // 要排除K个数
	if totalLen%2 == 0 {
		K--
	}

	mid1, mid2 := getNumberK(nums1, 0, nums2, 0, K)

	if totalLen%2 == 1 {
		return math.Min(mid1, mid2)
	}

	return (mid1 + mid2) / 2.0
}

func getNumberK(nums1 []int, start1 int, nums2 []int, start2 int, k int) (float64, float64) {
	if len(nums1) == 0 || start1 >= len(nums1) {
		pos := k + start2
		if pos < len(nums2) {
			return float64(nums2[pos]), float64(nums2[min(len(nums2)-1, pos+1)])
		}
		return math.MaxFloat64, math.MaxFloat64
	}
	if len(nums2) == 0 || start2 >= len(nums2) {
		pos := k + start1
		if pos < len(nums1) {
			return float64(nums1[pos]), float64(nums1[min(len(nums1)-1, pos+1)])
		}
		return math.MaxFloat64, math.MaxFloat64
	}

	if k <= 0 {
		index1 := min(start1, len(nums1)-1)
		index2 := min(start2, len(nums2)-1)
		m := float64(nums1[index1])
		n := float64(nums2[index2])
		if m < n {
			var mNext float64
			if (index1 + 1) < len(nums1) {
				mNext = float64(nums1[index1+1])
			} else {
				mNext = n
			}
			return m, math.Min(mNext, n)
		} else if m > n {
			var mNext float64
			if (index2 + 1) < len(nums2) {
				mNext = float64(nums2[index2+1])
			} else {
				mNext = m
			}
			return n, math.Min(mNext, m)
		} else {
			return m, n
		}
	}

	offset := max(k/2-1, 0)
	index1 := start1 + offset
	index2 := start2 + offset

	var m int
	var n int
	if index1 < len(nums1) {
		m = nums1[index1]
	} else {
		m = math.MaxInt32
	}
	if index2 < len(nums2) {
		n = nums2[index2]
	} else {
		n = math.MaxInt32
	}

	if m <= n {
		return getNumberK(nums1, start1+offset+1, nums2, start2, k-offset-1)
	}
	return getNumberK(nums1, start1, nums2, start2+offset+1, k-offset-1)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

/*
2020-10-18 [AC]
2020-11-02 [Review-001]
*/
