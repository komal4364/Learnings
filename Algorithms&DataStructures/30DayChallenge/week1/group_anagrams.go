/*
	https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/528/week-1/3288/

	Given an array of strings, group anagrams together.

	Example:

	Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
	Output:
	[
	["ate","eat","tea"],
	["nat","tan"],
	["bat"]
	]
	Note:

	All inputs will be in lowercase.
	The order of your output does not matter.
*/

package main

import "sort"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Len() int {
	return len(s)
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func hashStr(str string) string {
	chars := []rune(str)
	sort.Sort(sortRunes(chars))
	return string(chars)
}

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, elem := range strs {
		m[hashStr(elem)] = append(m[hashStr(elem)], elem)
	}
	out := [][]string{}
	for _, val := range m {
		out = append(out, val)
	}
	return out
}
