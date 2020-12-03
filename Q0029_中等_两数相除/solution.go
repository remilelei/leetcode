package Q0029_中等_两数相除

import "math"

/*
29. 两数相除
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2



示例 1:

输入: dividend = 10, divisor = 3
输出: 3
解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
解释: 7/-3 = truncate(-2.33333..) = -2


提示：

被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。本题中，如果除法结果溢出，则返回 2^31 − 1。
*/

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend == divisor {
		return 1
	}

	ret := 0
	isDiffSign := false
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		isDiffSign = true
	}
	if !isDiffSign {
		if (dividend > 0 && divisor > dividend) || (dividend < 0 && divisor < dividend) {
			return 0
		}
	}

	temp := 0
	for {
		nextTemp := temp
		if nextTemp == 0 {
			if isDiffSign {
				nextTemp = 0 - divisor
			} else {
				nextTemp = divisor
			}
		} else {
			nextTemp += temp
		}

		if (dividend < 0 && nextTemp < dividend) || (dividend >= 0 && nextTemp > dividend) {
			break
		}

		if ret == 0 {
			ret++
		} else {
			ret += ret
		}

		temp = nextTemp
	}

	if isDiffSign {
		ret = 0 - ret
	}

	if temp != 0 {
		ret = ret + divide(dividend-temp, divisor)
	}

	if ret > math.MaxInt32 || ret < math.MinInt32 {
		ret = math.MaxInt32
	}

	return ret
}
