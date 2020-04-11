/*
	https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/528/week-1/3286/

	Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

	Example:

	Input: [0,1,0,3,12]
	Output: [1,3,12,0,0]
	Note:

	You must do this in-place without making a copy of the array.
	Minimize the total number of operations.
*/

package main

func moveZeroes(nums []int) {
	for i, j := 0, 1; j < len(nums); {
		if nums[i] == 0 && nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i = i + 1
		} else if nums[i] != 0 {
			i = i + 1
		}
		j = j + 1
	}
}
