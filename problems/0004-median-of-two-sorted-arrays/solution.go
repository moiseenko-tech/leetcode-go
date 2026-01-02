// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p0004

// Problem: Median of Two Sorted Arrays
// URL: https://leetcode.com/problems/median-of-two-sorted-arrays/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// size, left and right pointers for nums1
	n1 := len(nums1)
	l1 := 0
	r1 := n1 - 1

	// size, left and right pointers for nums2
	n2 := len(nums2)
	l2 := 0
	r2 := n2 - 1

	n := n1 + n2

	// On each iteration both MIN and MAX elements will be eliminated.
	for n > 2 {
		// Eliminate MIN element.
		if l1 > r1 {
			// nums1 is empty
			l2++
		} else if l2 > r2 {
			// nums2 is empty
			l1++
		} else {
			if nums1[l1] < nums2[l2] {
				l1++
			} else {
				l2++
			}
		}
		n--

		// Eliminate MAX element.
		if l1 > r1 {
			// nums1 is empty
			r2--
		} else if l2 > r2 {
			// nums2 is empty
			r1--
		} else {
			if nums1[r1] > nums2[r2] {
				r1--
			} else {
				r2--
			}
		}
		n--
	}

	// Only 1 or 2 elements remain. Find the median from these elements.
	if l1 > r1 {
		// nums1 is empty
		if n == 1 {
			return float64(nums2[l2])
		}
		return float64(nums2[l2]+nums2[r2]) / 2
	} else if l2 > r2 {
		// nums2 is empty
		if n == 1 {
			return float64(nums1[l1])
		}
		return float64(nums1[l1]+nums1[r1]) / 2
	}
	return float64(nums1[l1]+nums2[l2]) / 2
}
