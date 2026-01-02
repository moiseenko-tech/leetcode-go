// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Andrei Moiseenko

package p0001

// Problem: Two Sum
// URL: https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		expectedValue := target - num

		expectedIndex, ok := seen[expectedValue]
		if ok {
			return []int{expectedIndex, i}
		}

		seen[num] = i
	}

	return []int{}
}
