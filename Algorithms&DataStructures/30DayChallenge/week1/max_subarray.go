/*
	https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/528/week-1/3285/

	Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

	Example:

	Input: [-2,1,-3,4,-1,2,1,-5,4],
	Output: 6
	Explanation: [4,-1,2,1] has the largest sum = 6.
	Follow up:

	If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/

package main

import "math"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	maxSum, currSum := -math.MaxInt64, 0
	for _, num := range nums {
		if currSum < 0 {
			currSum = 0
		}
		currSum = max(currSum+num, num)
		if currSum > maxSum {
			maxSum = currSum
		}
	}
	return maxSum
}
