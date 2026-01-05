// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p0448

// Problem: Find All Numbers Disappeared in an Array
// URL: https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

func findDisappearedNumbers(nums []int) []int {
	// TODO(amoiseenko): Improve performance.
	n := len(nums)

	seen := make([]bool, n)
	for _, v := range nums {
		seen[v-1] = true
	}

	result := make([]int, 0)
	for i := range n {
		if !seen[i] {
			result = append(result, i+1)
		}
	}

	return result
}
