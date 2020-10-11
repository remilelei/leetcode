package question1

/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Q1TwoSum : leetcode question 1 -- 两数之和
func Q1TwoSum(nums []int, target int) []int {
	resIndex1 := -1
	resIndex2 := -1
	indexSet := make(map[int]int)
	for index, value := range nums {
		factor1 := value
		factor2 := target - factor1
		index2, isExist := indexSet[factor2]
		if isExist && index2 != index {
			resIndex1 = index2
			resIndex2 = index
			break
		}
		indexSet[value] = index
	}

	return []int{resIndex1, resIndex2}
}
