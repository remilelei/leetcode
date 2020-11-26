package Q0014_简单_最长公共前缀

/*
14. 最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。
*/

func getCommonPrefix(a string, b string) string {
	if len(b) > len(a) {
		return getCommonPrefix(b, a)
	}
	commonDigit := 0
	for commonDigit < len(b) {
		if a[commonDigit] == b[commonDigit] {
			commonDigit++
		} else {
			break
		}
	}
	return b[:commonDigit]
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	curLongestCommonPrefix := strs[0]
	for i := 1; i < len(strs); i++ {
		curLongestCommonPrefix = getCommonPrefix(curLongestCommonPrefix, strs[i])
	}

	return curLongestCommonPrefix
}
