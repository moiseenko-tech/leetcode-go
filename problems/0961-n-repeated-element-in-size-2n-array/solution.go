// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p0961

// Problem: N-Repeated Element in Size 2N Array
// URL: https://leetcode.com/problems/n-repeated-element-in-size-2n-array/

func repeatedNTimes(nums []int) int {
	// The array contains N+1 distinct values, where one value is repeated N times.
	// Since the total length is 2N, the remaining N values appear exactly once.
	// Therefore, the element that occurs more than once is the target.
	seen := make(map[int]bool)

	for _, x := range nums {
		if seen[x] {
			return x
		}
		seen[x] = true
	}

	return 0
}
