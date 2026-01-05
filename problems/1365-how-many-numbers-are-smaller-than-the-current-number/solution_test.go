// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p1365

import (
	"reflect"
	"testing"
)

var tests = map[string]struct {
	nums   []int
	result []int
}{
	"8-1-2-2-3": {
		nums:   []int{8, 1, 2, 2, 3},
		result: []int{2, 1, 0, 3},
	},

	"7-7-7-7": {
		nums:   []int{7, 7, 7, 7},
		result: []int{0, 0, 0, 0},
	},
}

func TestSmallerNumbersThanCurrent(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := smallerNumbersThanCurrent(test.nums)
			expected := test.result
			if !reflect.DeepEqual(got, expected) {
				t.Fatalf("smallerNumbersThanCurrent(nums=%v) returned: %v; expected: %v",
					test.nums, got, expected)
			}
		})
	}
}
