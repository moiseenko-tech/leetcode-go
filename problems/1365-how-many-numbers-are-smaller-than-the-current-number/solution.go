// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p1365

// Problem: How Many Numbers Are Smaller Than the Current Number
// URL: https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/

func smallerNumbersThanCurrent(nums []int) []int {
	var maxNum int = 100
	n := len(nums)

	freq := make([]int, maxNum+1)
	for _, num := range nums {
		freq[num]++
	}

	/*
		nums:   8, 1, 2, 2, 3
		result: 4, 0, 1, 1, 3

		v = nums[i]             // nums[0] = 8
		result[i] = lessThan[v] // result[0] = lessThan[8]

		values:   0, 1, 2, 3, 4, 5, 6, 7, 8, 9
		freq:     0, 1, 2, 1, 0, 0, 0, 0, 1, 0
		lessThan: 0, 0, 1, 3, 4, 4, 4, 4, 4, 5

		lessThan[1] = lessThan[0] + freq[0]
	*/

	lessThan := make([]int, maxNum+1)
	for v := 1; v < maxNum+1; v++ {
		lessThan[v] = lessThan[v-1] + freq[v-1]
	}

	result := make([]int, n)
	for i := range n {
		result[i] = lessThan[nums[i]]
	}

	return result
}
