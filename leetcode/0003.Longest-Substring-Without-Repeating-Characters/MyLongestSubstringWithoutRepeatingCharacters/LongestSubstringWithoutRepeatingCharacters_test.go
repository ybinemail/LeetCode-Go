package MyLongestSubstringWithoutRepeatingCharacters

import "testing"

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {

	t.Log("")
	longstrings := []string{}
	longstrings = append(longstrings, "abba", "abcabcbb", "bbbbb", "pwwkew", "agecdg", "geeeegeg", "", "ckilbkd", "ckilbkdg")

	for _, v := range longstrings {
		t.Logf("The string %s has longest substring is: % d", v, lengthOfLongestSubstring2(v))
	}
}
