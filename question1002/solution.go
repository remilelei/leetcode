package question1002

import "math"

/*
给定仅有小写字母组成的字符串数组 A，返回列表中的每个字符串中都显示的全部字符（包括重复字符）组成的列表。例如，如果一个字符在每个字符串中出现 3 次，
但不是 4 次，则需要在最终答案中包含该字符 3 次。

你可以按任意顺序返回答案。



示例 1：

输入：["bella","label","roller"]
输出：["e","l","l"]
示例 2：

输入：["cool","lock","cook"]
输出：["c","o"]


提示：

1 <= A.length <= 100
1 <= A[i].length <= 100
A[i][j] 是小写字母


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-common-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func commonChars(A []string) []string {
	litterMap := make(map[int32]int)
	for stringIndex, val := range A {
		tempLitterMap := make(map[int32]int)
		for _, litter := range val {
			tempLitterMap[litter]++
		}
		if stringIndex == 0 {
			litterMap = tempLitterMap
		} else {
			for litter := range litterMap {
				tempVal, exist := tempLitterMap[litter]
				if exist {
					if tempVal < litterMap[litter] {
						litterMap[litter] = tempVal
					}
				} else {
					delete(litterMap, litter)
				}
			}
		}
	}

	var res []string
	for litter := range litterMap {
		for i := 0; i < litterMap[litter]; i++ {
			res = append(res, string(litter))
		}
	}
	return res
}

func commonCharsOffical(a []string) (ans []string) {
	minFreq := [26]int{}
	for i := range minFreq {
		minFreq[i] = math.MaxInt64
	}
	for _, word := range a {
		freq := [26]int{}
		for _, b := range word {
			freq[b-'a']++
		}
		for i, f := range freq[:] {
			if f < minFreq[i] {
				minFreq[i] = f
			}
		}
	}
	for i := byte(0); i < 26; i++ {
		for j := 0; j < minFreq[i]; j++ {
			ans = append(ans, string('a'+i))
		}
	}
	return
}
