package MyValidParentheses

import "strings"

func isValid(s string) bool {
	strSlice := strings.Split(s, "")
	var stack []string
	for _, v := range strSlice {
		isCouple := false
		var lastSign string
		if len(stack) >0{
			lastSign = stack[len(stack)-1]
		}
		switch v {
		case ")":
			if lastSign == "(" {
				isCouple = true
			}
		case "]":
			if lastSign == "[" {
				isCouple = true
			}
		case "}":
			if lastSign == "{" {
				isCouple = true
			}
		}
		if isCouple {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
