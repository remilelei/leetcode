package Q0017_中等_电话号码的字母组合

/*
17. 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。



示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
*/

var letters = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	combinations := []string{}
	for i := 0; i < len(digits); i++ {
		combinations = solution(digits, i, combinations[:])
	}
	return combinations
}

func solution(digits string, pos int, combinations []string) []string {
	letter := letters[digits[pos]-'2']
	if len(combinations) == 0 {
		for j := 0; j < len(letter); j++ {
			newCombination := string(letter[j])
			combinations = append(combinations, newCombination)
		}
	} else {
		l := len(combinations)
		for i := 0; i < l; i++ {
			combination := combinations[i]
			for j := 0; j < len(letter); j++ {
				newCombination := combination + string(letter[j])
				if j == 0 {
					combinations[i] = newCombination
				} else {
					combinations = append(combinations, newCombination)
				}
			}
		}
	}
	return combinations
}
