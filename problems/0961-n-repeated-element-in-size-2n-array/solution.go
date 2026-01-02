// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p0961

// Problem: N-Repeated Element in Size 2N Array
// URL: https://leetcode.com/problems/n-repeated-element-in-size-2n-array/

func repeatedNTimes(nums []int) int {
	// The array contains N+1 distinct values, where one value is repeated N times.
	// Since the total length is 2N, the remaining N values appear exactly once.
	// Therefore, the element that occurs more than once is the target.

	// For N >= 3, it is guaranteed that among all consecutive triples of elements
	// in the array, there exists at least one triple that contains a repeated value
	// (i.e., two equal numbers, regardless of their order inside the triple).
	// This follows from the fact that one value appears N times in an array of length 2N,

	// For the special case N = 2, this guarantee does not hold.
	// A counterexample is the array: [1, 2, 3, 1], where no consecutive triple
	// contains a duplicate. Therefore, for N = 2, all possible pairs of values
	// must be checked explicitly.

	// This property allows an O(1) memory solution by scanning the array using
	// consecutive triples of elements.

	if len(nums) == 4 {
		if nums[0] == nums[1] || nums[0] == nums[2] || nums[0] == nums[3] {
			return nums[0]
		}
		if nums[1] == nums[2] || nums[1] == nums[3] {
			return nums[1]
		}
		// nums[2] == nums[3]
		return nums[2]
	}

	for i := 0; i <= len(nums)-3; i++ {
		if nums[i] == nums[i+1] || nums[i] == nums[i+2] {
			return nums[i]
		}
		if nums[i+1] == nums[i+2] {
			return nums[i+1]
		}
	}

	return 0
}
