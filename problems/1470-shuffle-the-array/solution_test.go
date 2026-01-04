// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p1470

import (
	"reflect"
	"testing"
)

var tests = map[string]struct {
	nums   []int
	n      int
	result []int
}{
	"2-5-1-3-4-7--3": {
		nums:   []int{2, 5, 1, 3, 4, 7},
		n:      3,
		result: []int{2, 3, 5, 4, 1, 7},
	},
	"1-2-3-4-4-3-2-1--4": {
		nums:   []int{1, 2, 3, 4, 4, 3, 2, 1},
		n:      4,
		result: []int{1, 4, 2, 3, 3, 2, 4, 1},
	},
	"1-1-2-2--2": {
		nums:   []int{1, 1, 2, 2},
		n:      2,
		result: []int{1, 2, 1, 2},
	},
}

func TestRepeatedNTimes(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := shuffle(test.nums, test.n)
			expected := test.result
			if !reflect.DeepEqual(got, expected) {
				t.Fatalf("shuffle(nums=%v, n=%v) returned: %v; expected: %v",
					test.nums, test.n, got, expected)
			}
		})
	}
}
