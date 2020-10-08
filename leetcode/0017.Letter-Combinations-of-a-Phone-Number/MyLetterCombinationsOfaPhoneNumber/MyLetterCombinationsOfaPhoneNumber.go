package mylettercombinationsofaphonenumber

import (
	"fmt"
	"strings"
)

func letterCombinations(digits string) []string {
	var resultArr []string
	var res string

	if len(digits) == 0 {
		return resultArr
	}
	numb := strings.Split(digits, "")
	numbDic := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	var numbSlice [][]string
	for _, v := range numb {
		if value, ok := numbDic[v]; ok {
			numbSlice = append(numbSlice, value)
		}
	}
	getString(numbSlice, &res, &resultArr)
	fmt.Printf("result: %v \n", resultArr)
	return resultArr
}

func getString(slice [][]string, res *string, resultArr *[]string) {
	length := len(slice)
	for _, v := range slice[0] {
		if length > 1 {
			*res += v
			getString(slice[1:length], res, resultArr)
		} else {
			*resultArr = append(*resultArr, *res+v)
		}
	}
	if len(*res) > 0 {
		*res = (*res)[0 : len(*res)-1]
	}
}
