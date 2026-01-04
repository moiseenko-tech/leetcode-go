// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p1929

// Problem: Concatenation of Array
// URL: https://leetcode.com/problems/concatenation-of-array/

func getConcatenation(nums []int) []int {
	n := len(nums)
	result := make([]int, 2*n)

	for i := range n {
		result[i] = nums[i]
		result[n+i] = nums[i]
	}

	return result
}
