// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p0448

import (
	"reflect"
	"testing"
)

var tests = map[string]struct {
	nums   []int
	result []int
}{
	"4-3-2-7-8-2-3-1": {
		nums:   []int{4, 3, 2, 7, 8, 2, 3, 1},
		result: []int{5, 6},
	},
	"1-1": {
		nums:   []int{1, 1},
		result: []int{2},
	},
}

func TestFindDisappearedNumbers(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := findDisappearedNumbers(test.nums)
			expected := test.result
			if !reflect.DeepEqual(got, expected) {
				t.Fatalf("findDisappearedNumbers(nums=%v) returned: %v; expected: %v",
					test.nums, got, expected)
			}
		})
	}
}
