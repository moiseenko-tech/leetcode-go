// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p0004

import (
	"testing"
)

var tests = map[string]struct {
	nums1  []int
	nums2  []int
	result float64
}{
	"1-3--2": {
		nums1:  []int{1, 3},
		nums2:  []int{2},
		result: 2.0,
	},
	"1-2--3-4": {
		nums1:  []int{1, 2},
		nums2:  []int{3, 4},
		result: 2.5,
	},
}

func TestFindMedianSortedArrays(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := findMedianSortedArrays(test.nums1, test.nums2)
			expected := test.result
			if got != expected {
				t.Fatalf("findMedianSortedArrays(nums1=%v, nums2=%v) returned: %v; expected: %v",
					test.nums1, test.nums2, got, expected)
			}
		})
	}
}
