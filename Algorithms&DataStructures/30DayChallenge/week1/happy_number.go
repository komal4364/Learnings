/*
	https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/528/week-1/3284/

	Write an algorithm to determine if a number n is "happy".

	A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.

	Return True if n is a happy number, and False if not.

	Example:

	Input: 19
	Output: true
	Explanation:
	12 + 92 = 82
	82 + 22 = 68
	62 + 82 = 100
	12 + 02 + 02 = 1
*/

package main

func getDigits(n int) []int {
	out := []int{}
	for n != 0 {
		out = append(out, n%10)
		n = n / 10
	}
	return out
}

func isHappyNum(n int, m map[int]bool) bool {
	if n == 1 {
		return true
	}
	digits := getDigits(n)
	out := 0
	for _, digit := range digits {
		out = out + digit*digit
	}
	if out == 1 {
		return true
	}
	_, exist := m[out]
	if exist == true {
		return false
	}
	m[out] = true
	return isHappyNum(out, m)
}
func isHappy(n int) bool {
	m := map[int]bool{}
	return isHappyNum(n, m)
}
