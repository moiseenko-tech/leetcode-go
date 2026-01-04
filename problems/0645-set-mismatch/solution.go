// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p0645

// Problem: Set Mismatch
// URL: https://leetcode.com/problems/set-mismatch/

func findErrorNums(nums []int) []int {
	// TODO(amoiseenko): Improve performance.

	seen := make(map[int]bool)
	var double int
	var absent int

	for _, x := range nums {
		if seen[x] {
			double = x
		}
		seen[x] = true
	}

	for i := 1; i <= len(nums); i++ {
		if !seen[i] {
			absent = i
			break
		}
	}

	return []int{double, absent}
}
