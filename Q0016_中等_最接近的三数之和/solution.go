package Q0016_中等_最接近的三数之和

import (
	"math"
	"sort"
)

/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。



示例：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。


提示：

3 <= nums.length <= 10^3
-10^3 <= nums[i] <= 10^3
-10^4 <= target <= 10^4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum-closest
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func getAbsoluteDiff(a int, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	cloestSum := 0
	minDiff := math.MaxInt32
	for tuple1 := 0; tuple1 < len(nums)-2; tuple1++ {
		tuple2 := tuple1 + 1
		tuple3 := len(nums) - 1
		for tuple2 < tuple3 {
			tempSum := nums[tuple1] + nums[tuple2] + nums[tuple3]
			tempDiff := getAbsoluteDiff(tempSum, target)
			if tempDiff < minDiff {
				minDiff = tempDiff
				cloestSum = tempSum
			}

			if tempSum > target {
				tuple3--
			} else {
				tuple2++
			}
		}
	}
	return cloestSum
}

/*
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for p := 0; p < len(nums)-2; p++ {
		for l, r := p+1, len(nums)-1; l < r; {
			if math.Abs(float64(nums[p] + nums[l] + nums[r] - target)) < math.Abs(float64(res - target)) {
				res = nums[p] + nums[l] + nums[r]
			}
			if nums[p] + nums[l] + nums[r] < target {
				l++
			} else if nums[p] + nums[l] + nums[r] > target {
				r--
			} else {
				return target
			}
		}
	}
	return res
}

*/
