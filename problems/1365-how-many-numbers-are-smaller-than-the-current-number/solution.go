// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p1365

// Problem: How Many Numbers Are Smaller Than the Current Number
// URL: https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/

func smallerNumbersThanCurrent(nums []int) []int {
	// TODO(amoiseenko): Improve performance.
	return naiveSmallerNumbersThanCurrent(nums)
}

func naiveSmallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	for curr := range n {
		for neigh := range n {
			if nums[neigh] < nums[curr] {
				result[curr]++
			}
		}
	}

	return result
}
