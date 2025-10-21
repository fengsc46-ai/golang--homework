package main

import "fmt"

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
*/

func isValid(s string) bool {
	runes := []rune{}

	// 遍历字符串
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			runes = append(runes, v)
			fmt.Println(string(runes))
		} else {
			// 如果遇到右括号开头直接返回false
			if len(runes) == 0 {
				return false
			}
			// 如果遇到右括号，则弹出栈顶元素
			top := runes[len(runes)-1]
			runes = runes[:len(runes)-1]
			if (v == ')' && top != '(') ||
				(v == '}' && top != '{') ||
				(v == ']' && top != '[') {
				return false
			}
		}
	}
	return len(runes) == 0
}

func main() {
	s := "()[]{}{"
	fmt.Println(isValid(s))
}
