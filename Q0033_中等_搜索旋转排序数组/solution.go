package Q0033_中等_搜索旋转排序数组

/*
33. 搜索旋转排序数组
升序排列的整数数组 nums 在预先未知的某个点上进行了旋转（例如， [0,1,2,4,5,6,7] 经旋转后可能变为 [4,5,6,7,0,1,2] ）。

请你在数组中搜索 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。



示例 1：

输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
示例 2：

输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
示例 3：

输入：nums = [1], target = 0
输出：-1


提示：

1 <= nums.length <= 5000
-10^4 <= nums[i] <= 10^4
nums 中的每个值都 独一无二
nums 肯定会在某个点上旋转
-10^4 <= target <= 10^4

*/

func search(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch(nums []int, target int, begin int, end int) int {
	if begin < 0 || end >= len(nums) {
		return -1
	}
	if begin == end {
		if nums[begin] == target {
			return begin
		}
		return -1
	} else if begin > end {
		return -1
	}

	posA, posB, posC := begin, (begin+end)/2, end
	a, b, c := nums[posA], nums[posB], nums[posC]

	if a == target {
		return posA
	}
	if b == target {
		return posB
	}
	if c == target {
		return posC
	}

	if a < c {
		// 这一段没有被折叠过，直接二分查找
		if target < b {
			return binarySearch(nums, target, begin, posB-1)
		}
		return binarySearch(nums, target, posB+1, end)
	} else if a == c {
		// 可能折叠过，也有可能是一段全等数组
		if b == a {
			// 卧槽全等数组
			return -1
		} else if b > a {
			if target > a && target < b {
				return binarySearch(nums, target, begin, posB-1)
			}
			return binarySearch(nums, target, posB+1, end)
		} else {
			// b < a
			if target > b && target < c {
				return binarySearch(nums, target, posB+1, end)
			}
			return binarySearch(nums, target, begin, posB-1)
		}
	} else {
		// a > c 一定折叠过
		if b == a {
			// 前半等，从后面找
			return binarySearch(nums, target, posB+1, end)
		} else if b > a {
			if target > a && target < b {
				return binarySearch(nums, target, begin, posB-1)
			}
			return binarySearch(nums, target, posB+1, end)
		} else {
			// b < a
			if target > b && target < c {
				return binarySearch(nums, target, posB+1, end)
			}
			return binarySearch(nums, target, begin, posB-1)
		}
	}
}
