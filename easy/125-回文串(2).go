package easy

import (
	"regexp"
	"strings"
)

// 方法二：
func isPalindromeSelf(s string) bool {
	p1, p2 := 0, len(s)-1
	for p1 < p2 {
		if !isValid(s[p1]) {
			p1++
			continue
		}
		if !isValid(s[p2]) {
			p2--
			continue
		}
		if !strings.EqualFold(string(s[p1]), string(s[p2])) {
			return false
		}
		p1++
		p2--
	}
	return true
}

// 验证一个字符串是不是回文字符串，忽略大小写，只考虑字母和数字
func verifyHuiWenString(str string) bool {
	// 去掉空格
	str = strings.ReplaceAll(str, " ", "")
	// 字母全转大写
	str = strings.ToUpper(str)
	// 去掉字母及数字外的元素
	str = remainStrAndNum(str)
	// 对比反转后的字符串两者是否相同
	return str == reverseStr(str)
}

func reverseStr(str string) string {
	newStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		newStr += string(str[i])
	}
	return newStr
}

func remainStrAndNum(str string) string {
	newStr := ""
	for i := 0; i <= len(str)-1; i++ {
		if str[i] >= '0' && str[i] <= '9' || str[i] >= 'a' && str[i] <= 'z' || str[i] >= 'A' && str[i] <= 'Z' {
			newStr += string(str[i])
		}
	}
	return newStr
}

// 方法1
func isPalindrome(s string) bool {

	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	processedString := reg.ReplaceAllString(s, "")
	processedString = strings.ToLower(processedString)
	for i := 0; i < len(processedString)/2; i++ {
		if processedString[i] != processedString[len(processedString)-1-i] {
			return false
		}
	}
	return true
}

// 方法二：
func isPalindromeV2(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if !isValid(s[i]) {
			i++
			continue
		}
		if !isValid(s[j]) {
			j--
			continue
		}
		if !strings.EqualFold(string(s[i]), string(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}

func isValid(a byte) bool {
	if (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}
