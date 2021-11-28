package easy

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。

func strIsValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}
	stack := make([]string, 0)
	rule := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	for _, v := range s {
		oneStr := string(v)
		if isLeft(oneStr) {
			stack = append(stack, oneStr)
		} else {
			if len(stack) == 0 {
				return false
			}
			if rule[stack[len(stack)-1]] != oneStr {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func isLeft(s string) bool {
	if s == "(" || s == "{" || s == "[" {
		return true
	}
	return false
}
