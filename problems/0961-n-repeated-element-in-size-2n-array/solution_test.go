// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p0961

import (
	"testing"
)

var tests = map[string]struct {
	nums   []int
	result int
}{
	"1-2-3-3": {
		nums:   []int{1, 2, 3, 3},
		result: 3,
	},
	"2-1-2-5-3-2": {
		nums:   []int{2, 1, 2, 5, 3, 2},
		result: 2,
	},
	"5-1-5-2-5-3-5-4": {
		nums:   []int{5, 1, 5, 2, 5, 3, 5, 4},
		result: 5,
	},
	"9-5-6-9": {
		nums:   []int{9, 5, 6, 9},
		result: 9,
	},
}

func TestRepeatedNTimes(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := repeatedNTimes(test.nums)
			expected := test.result
			if got != expected {
				t.Fatalf("repeatedNTimes(nums=%v) returned: %v; expected: %v",
					test.nums, got, expected)
			}
		})
	}
}

func TestRepeatedNTimesWithMap(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := repeatedNTimesWithMap(test.nums)
			expected := test.result
			if got != expected {
				t.Fatalf("repeatedNTimesWithMap(nums=%v) returned: %v; expected: %v",
					test.nums, got, expected)
			}
		})
	}
}
