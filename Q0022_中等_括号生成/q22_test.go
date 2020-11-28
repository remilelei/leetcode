package Q0022_中等_括号生成

import "testing"

func TestQuestion22_1(t *testing.T) {
	res := generateParenthesis(3)
	t.Logf("res = %v", res)
}
