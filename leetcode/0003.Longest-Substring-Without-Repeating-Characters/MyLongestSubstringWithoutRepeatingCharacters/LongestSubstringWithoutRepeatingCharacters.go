package MyLongestSubstringWithoutRepeatingCharacters

import (
	"sort"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	// 考虑采用map进行判断是否存在，但是无法判定存在位置，bad case：abba
	arr := strings.Split(s, "")
	longest := []int{}
	count := 0
	sub := make(map[string]int)
	for i, v := range arr {
		k, ok := sub[v]
		if ok {
			count = i - k
		} else {
			count++
		}
		longest = append(longest, count)
		sub[v] = i
	}
	sort.Ints(longest)
	if len(longest) > 0 {
		return longest[len(longest)-1]
	}
	return 0
}

func lengthOfLongestSubstring2(s string) int {
	// 暴力拆分字符
	longest := []int{}
	for i := 0; i < len(s); i++ {
		length := 1
		for j := i + 1; j < len(s); j++ {
			// 双重循环判断是否存在
			if strings.ContainsAny(s[i:j], string(s[j])) {
				break
			}
			length++
		}
		longest = append(longest, length)
	}
	// 排序获取最长量
	sort.Ints(longest)
	if len(longest) > 0 {
		return longest[len(longest)-1]
	}
	return 0
}
