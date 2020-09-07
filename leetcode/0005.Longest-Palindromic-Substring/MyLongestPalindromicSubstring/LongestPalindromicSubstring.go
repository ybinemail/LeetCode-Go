package MyLongestPalindromicSubstring

import (
	"strings"
)

// 字符比较法 无法定位aaa,aaaa叠字位置
func longestPalindrome(s string) string {
	strarr := strings.Split(s, "")
	longestarr := []string{}

	for i := 0; i < len(strarr); i++ {
		//abcdcb abcddcbf abcdddcbf aaabaaaa
		char := strarr[i]
		if i == 0 {
			longestarr = append(longestarr, char)
		}
		// 设置每步跳的步长
		step := 0
		// 针对 abcddcbf中有，叠字符 dd ，设置叠字符长度重复长度
		same := 0
		for j := i - 1; j >= 0; j-- {
			step = i - j - same
			// 1.添加叠词判断 需要解决bananas => anana abcddcbf
			// 2.j == i-1 ,保证只在初始的时候，定位是否已叠词为中心
			if char == strarr[j] && j == i-1 {
				//出现叠词
				same++
				//单独叠词也计算回文
				longestarr = append(longestarr, strings.Join(strarr[j:i+1], ""))
			} else {
				//不相等，以当前词为中心，两边拓展查找
				if i+step < len(strarr) && strarr[j] == strarr[i+step] {
					longestarr = append(longestarr, strings.Join(strarr[j:i+step+1], ""))
				} else {
					break
				}
			}
		}

	}
	maxstring := ""
	for _, v := range longestarr {
		if len(maxstring) < len(v) {
			maxstring = v
		}
	}
	return maxstring
}

// 字符占位,针对大字符串，占用时间较长
func longestPalindrome2(s string) string {
	arr := strings.Split(s, "")
	// 1.把arr按照某中规则，重新合并到arrStr中
	arrStr := []string{}
	for i, v := range arr {
		arrStr = append(arrStr, v)
		// 2.重头不开始，排除最后一个，如果当前位和下一位相同，采用字符进行占位
		if i < len(arr)-1 && v == arr[i+1] {
			arrStr = append(arrStr, "#")
		}
	}
	maxstr := ""
	hasCount := 0
	for i := 0; i < len(arrStr); i++ {
		//abcdcb abcddcbf abcdddcbf aaabaaaa
		if i == 0 {
			maxstr = arrStr[0]
		}
		if arrStr[i] == "#" {
			hasCount++
		}
		// 设置每步跳的步长
		step := 0
		for j := i - 1; j >= 0; j-- {
			step = i - j
			//以当前词为中心，两边拓展查找
			if i+step < len(arrStr) && arrStr[j] == arrStr[i+step] {
				strlen := 0
				if arrStr[i] == "#" {
					strlen = step*2 + 1 - hasCount
				} else {
					strlen = step*2 + 1 - hasCount - 1
				}
				if len(maxstr) <= strlen {
					maxstr = strings.Join(arrStr[j:i+step+1], "")
				}
			} else {
				break
			}
		}
	}
	return strings.ReplaceAll(maxstr, "#", "")
}

// 定位字符中心，分别已当前字符及当前及下一个字符为中心
func longestPalindrome3(s string) string {
	maxstr := ""
	for i, v := range s {
		if i == 0 {
			maxstr = string(v)
		}
		for j := i + 1; j < len(s); j++ {
			step := j - i
			//以当前词为中心，两边拓展查找
			if i-step >= 0 && s[j] == s[i-step] {
				if len(maxstr) <= (step*2 + 1) {
					maxstr = s[i-step : j+1]
				}
			} else {
				break
			}
		}

		for k := i + 1; k < len(s); k++ {
			//以当前字符和下一个字符为中心
			sel := k - i - 1
			if i-sel >= 0 && s[k] == s[i-sel] {
				if len(maxstr) <= (sel*2 + 1) {
					maxstr = s[i-sel : k+1]
				}
			} else {
				break
			}
		}
	}
	return maxstr
}
