// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p0645

import (
	"reflect"
	"testing"
)

var tests = map[string]struct {
	nums   []int
	result []int
}{
	"1-2-2-4": {
		nums:   []int{1, 2, 2, 4},
		result: []int{2, 3},
	},
	"1-1": {
		nums:   []int{1, 1},
		result: []int{1, 2},
	},
}

func TestRepeatedNTimes(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := findErrorNums(test.nums)
			expected := test.result
			if !reflect.DeepEqual(got, expected) {
				t.Fatalf("findErrorNums(nums=%v) returned: %v; expected: %v",
					test.nums, got, expected)
			}
		})
	}
}
