// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Andrei Moiseenko

package p0167

// Problem: Two Sum II - Input Array Is Sorted
// URL: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		}

		if sum < target {
			l++
		} else {
			r--
		}
	}

	return []int{}
}
