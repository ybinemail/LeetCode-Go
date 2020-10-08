package mylettercombinationsofaphonenumber

import (
	"reflect"
	"testing"
)

/*给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。*/

// 考虑 23，223,299,99
func TestLetterCombinations(t *testing.T) {
	type test struct {
		input string
		want  []string
	}
	tests := map[string]test{
		"case_1": {
			"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		"case_2": {
			"234", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		"case_3": {
			"", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
	}

	for name, ts := range tests {
		t.Run(name, func(t *testing.T) {
			got := letterCombinations(ts.input)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("name:%s excepted:%#v, got:%#v", name, ts.want, got)
			}
		})
	}
}
