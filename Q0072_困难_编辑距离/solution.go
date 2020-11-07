package Q0072_困难_编辑距离

import "fmt"

/*
给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符


示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：

输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')


提示：

0 <= word1.length, word2.length <= 500
word1 和 word2 由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/edit-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minDistance(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)

	if len1 == 0 {
		return len2
	} else if len2 == 0 {
		return len1
	}

	// 1. make array
	dp := make([][]int, len1) //[len1][len2]int

	/*
	   2. define array element relation
	   f(x, y) = min( f(x-1,y), f(x-1,y-1), f(x,y-1) ) + (word1[x]!=word2[y])
	*/

	// 3. init array
	min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	safeRead := func(x, y int) int {
		if x < 0 || x >= len1 || y < 0 || y >= len2 {
			return 500 // 本题中不会大于500
		}
		return dp[x][y]
	}
	for i := 0; i < len1; i++ {
		dp[i] = make([]int, len2)
		bit := 0
		if word1[i] != word2[0] {
			bit = 1
		}
		dp[i][0] = min(i+bit, safeRead(i-1, 0)+1)
	}
	for i := 0; i < len2; i++ {
		bit := 0
		if word1[0] != word2[i] {
			bit = 1
		}
		dp[0][i] = min(i+bit, safeRead(0, i-1)+1)
	}

	// 4. build array
	for i := 1; i < len1; i++ {
		for j := 1; j < len2; j++ {
			// bit := 0
			offset := 0
			if i > j {
				offset = i - j
			} else {
				offset = j - i
			}
			// if word1[i] != word2[j] {
			// 	bit = 1
			// }
			// dp[i][j] = min(min(safeRead(i-1, j), safeRead(i, j-1)), safeRead(i-1, j-1)) + bit
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
			if offset > dp[i][j] {
				dp[i][j] = offset
			}
		}
	}

	for _, v := range dp {
		fmt.Println(v)
	}

	return dp[len1-1][len2-1]
}

func minDistance_Awesome(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	if l1 == 0 || l2 == 0 {
		return l1 + l2
	}
	dp := make([]int, l2+1)
	for i := 0; i < l2+1; i++ {
		dp[i] = i
	}
	var min = func(x, y, z int) int {
		if x <= y && x <= z {
			return x
		}
		if y <= z {
			return y
		}
		return z
	}
	for i := 1; i < l1+1; i++ {
		preIJ := dp[0]
		dp[0] = i
		for j := 1; j < l2+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[j], preIJ = min(dp[j-1]+1, dp[j]+1, preIJ), dp[j]
			} else {
				dp[j], preIJ = min(dp[j-1]+1, dp[j]+1, preIJ+1), dp[j]
			}
		}
	}
	return dp[l2]
}
