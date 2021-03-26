package main

import (
	"fmt"
)

/*
Given a string s, find the length of the longest substring without repeating characters.

EXAMPLE
	Input: s = "pwwkew"
	Output: 3
	Explanation: The answer is "wke", with the length of 3.
	Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

	Input: s = ""
	Output: 0

CONSTRAINTS
	0 <= s.length <= 5 * 104
	s consists of English letters, digits, symbols and spaces.
*/

func lengthOfLongestSubstring(s string) int {
	result := 0
	m := make(map[uint8]int)
	i := 0
	for j := range s {
		if v, ok := m[s[j]]; ok {
			i = max(i, v)
		}
		result = max(result, j - i + 1)
		m[s[j]] = j + 1
	}
	return result
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	i := lengthOfLongestSubstring("pwwkew")
	fmt.Print(i)
}
