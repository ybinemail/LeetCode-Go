package longestvalidparentheses

import "strings"

func longestValidParentheses(s string) int {
	var stack []string
	slice := strings.Split(s, "")
	longestNum := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] != "(" {
			continue
		}
		for j := i + 1; j < len(slice); j++ {
			if slice[j] != ")" {
				continue
			}
			sli := slice[i : j+1]
			if isValidParentheses(&sli, &stack) && longestNum < (j-i+1) {
				longestNum = j - i + 1
			}
		}
	}
	return longestNum
}

func longestValidParentheses2(s string) int {
	var stack []string
	slice := strings.Split(s, "")
	longestNum := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] != "(" {
			continue
		}
		for j := i + 1; j < len(slice); j++ {
			if slice[j] != ")" {
				continue
			}
			sli := slice[i : j+1]
			if isValidParentheses(&sli, &stack) && longestNum < (j-i+1) {
				longestNum = j - i + 1
			}
		}
	}
	return longestNum
}

func isValidParentheses(sli *[]string, stack *[]string) bool {
	for _, v := range *sli {
		if v == ")" {
			if len(*stack) > 0 && (*stack)[len(*stack)-1] == "(" {
				*stack = (*stack)[:len(*stack)-1]
			} else {
				//*stack = append(*stack, v) 直接判断
				return false
			}
		} else {
			*stack = append(*stack, v)
		}
	}
	if len(*stack) == 0 {
		return true
	}
	*stack = (*stack)[0:0]
	return false
}
