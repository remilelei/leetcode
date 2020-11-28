package Q0022_中等_括号生成

/*
22. 括号生成
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。



示例：

输入：n = 3
输出：[
       "((()))",
       "(()())",
       "(())()",
       "()(())",
       "()()()"
     ]
*/

func generateParenthesis(n int) []string {
	ret := []string{}

	if n <= 0 {
		return ret
	}

	if n == 1 {
		return []string{"()"}
	}
	//(n%2 == 1 && i <= n/2) || (n%2 == 0 && i < n/2)
	for i := 0; i < n; i++ {
		mergeRet1 := generateParenthesis(i)
		mergeRet2 := generateParenthesis(n - 1 - i)
		if len(mergeRet1) > 0 {
			for _, v1 := range mergeRet1 {
				if len(mergeRet2) > 0 {
					for _, v2 := range mergeRet2 {
						ret = append(ret, "("+v1+")"+v2)
						// ret = append(ret, v1+"("+v2+")")
					}
				} else {
					// ret = append(ret, "()"+v1)
					ret = append(ret, "("+v1+")")
				}
			}
		} else {
			if len(mergeRet2) > 0 {
				for _, v2 := range mergeRet2 {
					ret = append(ret, "()"+v2)
					// ret = append(ret, "("+v2+")")
				}
			}
		}
	}
	return ret
}
