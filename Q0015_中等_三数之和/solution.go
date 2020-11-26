package Q0015_中等_三数之和

import (
	"fmt"
	"math"
	"sort"
)

/*
15. 三数之和
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

*/

func threeSum(nums []int) [][]int {
	ret := [][]int{}

	sort.Ints(nums)

	lastTuple1 := math.MinInt32
	for i := 0; i < len(nums); i++ {
		tuple1 := nums[i]
		if lastTuple1 == tuple1 {
			continue
		}
		lastTuple1 = tuple1
		left := i + 1
		right := len(nums) - 1
		lastTuple2 := math.MinInt32
		for left < right {
			tuple2 := nums[left]
			tuple3 := nums[right]
			sum := tuple1 + tuple2 + tuple3
			if sum == 0 {
				if lastTuple2 == tuple2 {
					left++
					continue
				}
				lastTuple2 = tuple2
				elem := []int{tuple1, tuple2, tuple3}
				ret = append(ret, elem)
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	fmt.Printf("%v", ret)
	return ret
}
