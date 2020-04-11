/*

	https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/528/week-1/3289/

	Given an integer array arr, count element x such that x + 1 is also in arr.

	If there're duplicates in arr, count them seperately.



	Example 1:

	Input: arr = [1,2,3]
	Output: 2
	Explanation: 1 and 2 are counted cause 2 and 3 are in arr.
	Example 2:

	Input: arr = [1,1,3,3,5,5,7,7]
	Output: 0
	Explanation: No numbers are counted, cause there's no 2, 4, 6, or 8 in arr.
	Example 3:

	Input: arr = [1,3,2,3,5,0]
	Output: 3
	Explanation: 0, 1 and 2 are counted cause 1, 2 and 3 are in arr.
	Example 4:

	Input: arr = [1,1,2,2]
	Output: 2
	Explanation: Two 1s are counted cause 2 is in arr.


	Constraints:

	1 <= arr.length <= 1000
	0 <= arr[i] <= 1000

*/

package main

import "sort"

func countElements(arr []int) int {
	sort.Ints(arr)
	m := map[int][]int{}
	for idx, elem := range arr {
		m[elem] = append(m[elem], idx)
	}

	count := 0
	for _, elem := range arr {
		next := m[elem+1]
		if len(next) == 0 {
			continue
		}
		count = count + 1
		m[elem] = m[elem][1:]
	}
	return count
}