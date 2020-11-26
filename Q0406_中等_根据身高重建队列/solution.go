package Q0406_中等_根据身高重建队列

import (
	"sort"
)

/**
406. 根据身高重建队列
假设有打乱顺序的一群人站成一个队列。 每个人由一个整数对 (h, k) 表示，其中 h 是这个人的身高，k 是应该排在这个人前面且身高大于或等于 h 的人数。 例如：[5,2] 表示前面应该有 2 个身高大于等于 5 的人，而 [5,0] 表示前面不应该存在身高大于等于 5 的人。

编写一个算法，根据每个人的身高 h 重建这个队列，使之满足每个整数对 (h, k) 中对人数 k 的要求。

示例：

输入：[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
输出：[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]


提示：

总人数少于 1100 人。
*/

func reconstructQueue(people [][]int) [][]int {
	if len(people) <= 0 {
		return people
	}

	// 第一次排序，按身高降序排列
	less := func(x int, y int) bool {
		if people[x][0] > people[y][0] {
			return true
		} else if people[x][0] == people[y][0] {
			return people[x][1] < people[y][1]
		}
		return false
	}
	sort.Slice(people, less)

	// 第二次排序，安排按照要求插队
	insert := func(current int, target int) {
		temp := people[current]
		for current != target && current > 0 {
			people[current] = people[current-1]
			current--
		}
		people[current] = temp
	}
	for i, v := range people {
		insert(i, v[1])
	}
	return people
}
