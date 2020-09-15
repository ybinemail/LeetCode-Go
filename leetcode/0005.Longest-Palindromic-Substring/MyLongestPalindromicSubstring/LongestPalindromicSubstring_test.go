package MyLongestPalindromicSubstring

import (
	"reflect"
	"testing"
)

func TestLongestPalindromicSubstring(t *testing.T) {
	type test struct {
		s    string
		want string
	}
	tests := map[string]test{
		"case_1":  {"aaabaaaa", "aaabaaa"},
		"case_2":  {"aaa", "aaa"},
		"case_3":  {"bbaa", "bb"},
		"case_4":  {"bba", "bb"},
		"case_5":  {"cbb", "bb"},
		"case_6":  {"cbbd", "bb"},
		"case_7":  {"aaaa", "aaaa"},
		"case_8":  {"222020221", "2202022"},
		"case_9":  {"abcddcbf", "bcddcb"},
		"case_10": {"aabaa", "aabaa"},
		"case_11": {"bananas", "bananas"},
		"case_12": {"", ""},
		"case_13": {"ab", "a"},
		"case_14": {"a", "a"},
		"case_15": {"abcdcbahihba", "abcdcba"},
	}

	for name, ts := range tests {
		got := longestPalindrome3(ts.s)
		if !reflect.DeepEqual(got, ts.want) {
			t.Errorf("name:%s excepted:%#v, got:%#v", name, ts.want, got)
		}
	}
}
