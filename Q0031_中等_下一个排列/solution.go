package Q0031_中等_下一个排列

import (
	"sort"
)

/*
31. 下一个排列
实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须 原地 修改，只允许使用额外常数空间。



示例 1：

输入：nums = [1,2,3]
输出：[1,3,2]
示例 2：

输入：nums = [3,2,1]
输出：[1,2,3]
示例 3：

输入：nums = [1,1,5]
输出：[1,5,1]
示例 4：

输入：nums = [1]
输出：[1]


提示：

1 <= nums.length <= 100
0 <= nums[i] <= 100
*/

func nextPermutation(nums []int) {
	// 1. 找升序对
	ascPairHead := -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			// 找到升序对
			ascPairHead = i - 1
			break
		}
	}
	if ascPairHead == -1 { // 没有升序对，已是最大字典序，直接返回升序排序串
		sort.Ints(nums)
		return
	}

	// 2. 从尾部找出比升序对头部大的第一个数
	target := nums[ascPairHead]
	targetIndex := -1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > target {
			target = nums[i]
			targetIndex = i
			break
		}
	}

	// 3. 将这个数放到升序对头部并维护后方升序
	nums[targetIndex] = nums[ascPairHead]
	nums[ascPairHead] = target
	for i, j := ascPairHead+1, len(nums)-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

}
