// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p0485

import (
	"testing"
)

var tests = map[string]struct {
	nums   []int
	result int
}{
	"1-1-0-1-1-1": {
		nums:   []int{1, 1, 0, 1, 1, 1},
		result: 3,
	},
	"1-0-1-1-0-1": {
		nums:   []int{1, 0, 1, 1, 0, 1},
		result: 2,
	},
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := findMaxConsecutiveOnes(test.nums)
			expected := test.result
			if got != expected {
				t.Fatalf("findMaxConsecutiveOnes(nums=%v) returned: %v; expected: %v",
					test.nums, got, expected)
			}
		})
	}
}
