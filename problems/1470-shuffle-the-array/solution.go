// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p1470

// Problem: Shuffle the Array
// URL: https://leetcode.com/problems/shuffle-the-array/

func shuffle(nums []int, n int) []int {
	// TODO(amoiseenko): Improve performance.

	result := make([]int, 2*n)

	for i := range n {
		result[i*2] = nums[i]
		result[i*2+1] = nums[n+i]
	}

	return result
}
