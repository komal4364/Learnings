/*
	https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/529/week-2/3291/

	Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character.

	Example 1:

	Input: S = "ab#c", T = "ad#c"
	Output: true
	Explanation: Both S and T become "ac".
	Example 2:

	Input: S = "ab##", T = "c#d#"
	Output: true
	Explanation: Both S and T become "".
	Example 3:

	Input: S = "a##c", T = "#a#c"
	Output: true
	Explanation: Both S and T become "c".
	Example 4:

	Input: S = "a#c", T = "b"
	Output: false
	Explanation: S becomes "c" while T becomes "b".
	Note:

	1 <= S.length <= 200
	1 <= T.length <= 200
	S and T only contain lowercase letters and '#' characters.
	Follow up:

	Can you solve it in O(N) time and O(1) space?
*/
package main

func backspaceCompare(S string, T string) bool {
	inp := []rune{}
	for _, char := range S {
		if char == '#' {
			if len(inp) == 0 {
				continue
			}
			inp = inp[:len(inp)-1]
		} else {
			inp = append(inp, char)
		}
	}
	tinp := []rune{}
	for _, char := range T {
		if char == '#' {
			if len(tinp) == 0 {
				continue
			}
			tinp = tinp[:len(tinp)-1]
		} else {
			tinp = append(tinp, char)
		}
	}
	return string(inp) == string(tinp)
}
