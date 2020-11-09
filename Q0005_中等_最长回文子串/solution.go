package Q0005_中等_最长回文子串

import (
	"bytes"
)

/*
5. 最长回文子串
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
*/

func longestPalindromeOld(s string) string {
	l := len(s)
	if l <= 1 {
		return s
	}

	dp := make([][]int, l)
	resI, resJ := -1, -1
	palindromeHalf := 0

	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
		if s[l-i-1] == s[0] {
			dp[i][0] = 1
			palindromeHalf = 1
			resI, resJ = i, 0
		}
	}
	for i := 0; i < l; i++ {
		if s[i] == s[l-1] {
			dp[0][i] = 1
			palindromeHalf = 1
			resI, resJ = 0, i
		}
	}

	for i := 1; i < l; i++ {
		for j := 1; j < l; j++ {
			if s[j] == s[l-i-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				isPalindrome := (i-j < 2*l) && (j+i < 2*dp[i-1][j-1]+l-1)
				if i+j >= l && dp[i][j] > palindromeHalf && isPalindrome {
					palindromeHalf = dp[i][j]
					resI = i
					resJ = j
				}
			}
		}
	}

	resLen := palindromeHalf
	b := make([]byte, resLen)
	i, j, k := resI, resJ, palindromeHalf-1
	for i >= 0 && j >= 0 && dp[i][j] > 0 {
		b[k] = s[j]
		i--
		j--
		k--
	}

	buf := new(bytes.Buffer)
	buf.Write(b)
	return buf.String()
}

func longestPalindrome(s string) string {
	maxPalindrome := ""
	maxPalindromeLen := 0

	curPos := 0
	for curPos < len(s) {
		left, right := curPos, curPos
		for right+1 < len(s) && s[right+1] == s[left] {
			right++
		}
		nextPos := right + 1
		for left > 0 && right < len(s)-1 && s[right+1] == s[left-1] {
			left--
			right++
		}
		if left < 0 {
			left = 0
		}
		if right >= len(s) {
			right = len(s) - 1
		}
		curLen := right - left + 1
		if curLen > maxPalindromeLen {
			maxPalindromeLen = curLen
			maxPalindrome = s[left : right+1]
		}
		curPos = nextPos
	}

	return maxPalindrome
}
