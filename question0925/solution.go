package question0925

/*
你的朋友正在使用键盘输入他的名字 name。偶尔，在键入字符 c 时，按键可能会被长按，而字符可能被输入 1 次或多次。
你将会检查键盘输入的字符 typed。如果它对应的可能是你的朋友的名字（其中一些字符可能被长按），那么就返回 True。

示例 1：
输入：name = "alex", typed = "aaleex"
输出：true
解释：'alex' 中的 'a' 和 'e' 被长按。

示例 2：
输入：name = "saeed", typed = "ssaaedd"
输出：false
解释：'e' 一定需要被键入两次，但在 typed 的输出中不是这样。

示例 3：
输入：name = "leelee", typed = "lleeelee"
输出：true

示例 4：
输入：name = "laiden", typed = "laiden"
输出：true
解释：长按名字中的字符并不是必要的。

提示：
name.length <= 1000
typed.length <= 1000
name 和 typed 的字符都是小写字母。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/long-pressed-name
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isLongPressedName(name string, typed string) bool {
	typedIndex, nameIndex := 0, 0
	for typedIndex < len(typed) {
		if nameIndex >= len(name) {
			nameIndex = len(name) - 1
		}
		if typed[typedIndex] == name[nameIndex] {
			// 若相同，同时进入下一步
			typedIndex++
			nameIndex++
		} else {
			//若不同，看看typed的字符是不是name中上一个字符的重复
			if nameIndex > 0 {
				if name[nameIndex-1] == typed[typedIndex] {
					for name[nameIndex] != typed[typedIndex] {
						if name[nameIndex-1] != typed[typedIndex] {
							return false
						}
						typedIndex++
						if typedIndex >= len(typed) {
							return false
						}
					}
				} else {
					return false
				}
			} else {
				// name没有上一个字符了，此时一定不可能是朋友的名字
				return false
			}
		}
	}

	if nameIndex < len(name) {
		return false
	}

	return true
}
