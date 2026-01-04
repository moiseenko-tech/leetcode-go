// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p0485

// Problem: Max Consecutive Ones
// URL: https://leetcode.com/problems/max-consecutive-ones/

func findMaxConsecutiveOnes(nums []int) int {
	curLen := 0
	maxLen := 0

	for _, x := range nums {
		if x != 1 {
			curLen = 0
			continue
		}
		curLen++
		if curLen > maxLen {
			maxLen = curLen
		}
	}

	return maxLen
}
