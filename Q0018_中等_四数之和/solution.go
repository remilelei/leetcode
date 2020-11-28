package Q0018_中等_四数之和

import "sort"

/*
18. 四数之和
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]

*/

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	length := len(nums)

	ret := [][]int{}
	comboo1 := 0
	for comboo1 < length-3 {
		comboo2 := comboo1 + 1
		for comboo2 < length-2 {
			comboo3 := comboo2 + 1
			comboo4 := length - 1
			for comboo3 < comboo4 {
				sum := nums[comboo1] + nums[comboo2] + nums[comboo3] + nums[comboo4]
				if sum == target {
					ret = append(ret, []int{nums[comboo1], nums[comboo2], nums[comboo3], nums[comboo4]})
				}

				comboo3Value := nums[comboo3]
				comboo4Value := nums[comboo4]
				if sum < target {
					for comboo3 < length && nums[comboo3] == comboo3Value {
						comboo3++
					}
				} else {
					for comboo4 >= 0 && nums[comboo4] == comboo4Value {
						comboo4--
					}
				}
			}

			comboo2Value := nums[comboo2]
			for comboo2 < length-2 && nums[comboo2] == comboo2Value {
				comboo2++
			}
		}
		comboo1Value := nums[comboo1]
		for comboo1 < length-3 && nums[comboo1] == comboo1Value {
			comboo1++
		}
	}

	return ret
}
