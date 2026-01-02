// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Andrei Moiseenko

package p0167

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
		result: []int{1, 2},
	},
	"2-3-4-6": {
		nums:   []int{2, 3, 4},
		target: 6,
		result: []int{1, 3},
	},
	"-1-0--1": {
		nums:   []int{-1, 0},
		target: -1,
		result: []int{1, 2},
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
