/*
	https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/528/week-1/3283/

	Given a non-empty array of integers, every element appears twice except for one. Find that single one.

	Note:

	Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

	Example 1:

	Input: [2,2,1]
	Output: 1
	Example 2:

	Input: [4,1,2,1,2]
	Output: 4
*/

package main

func singleNumber(nums []int) int {
    out := 0
    for _, val := range nums {
        out = out ^ val
    }
    return out
}