// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Andrei Moiseenko

package p0001

import (
	"reflect"
	"testing"
)

var tests = map[string]struct {
	nums   []int
	target int
	result []int
}{
	"2-7-11-15-9": {
		nums:   []int{2, 7, 11, 15},
		target: 9,
		result: []int{0, 1},
	},
	"3-2-4-6": {
		nums:   []int{3, 2, 4},
		target: 6,
		result: []int{1, 2},
	},
	"3-3-6": {
		nums:   []int{3, 3},
		target: 6,
		result: []int{0, 1},
	},
}

func TestTwoSum(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := twoSum(test.nums, test.target)
			expected := test.result
			if !reflect.DeepEqual(got, expected) {
				t.Fatalf("twoSum(nums=%v, target=%v) returned: %v; expected: %v",
					test.nums, test.target, got, expected)
			}
		})
	}
}
