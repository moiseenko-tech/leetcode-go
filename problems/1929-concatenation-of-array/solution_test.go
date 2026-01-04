// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p1929

import (
	"reflect"
	"testing"
)

var tests = map[string]struct {
	nums   []int
	result []int
}{
	"1-2-1": {
		nums:   []int{1, 2, 1},
		result: []int{1, 2, 1, 1, 2, 1},
	},
	"1-3-2-1": {
		nums:   []int{1, 3, 2, 1},
		result: []int{1, 3, 2, 1, 1, 3, 2, 1},
	},
}

func TestGetConcatenation(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := getConcatenation(test.nums)
			expected := test.result
			if !reflect.DeepEqual(got, expected) {
				t.Fatalf("getConcatenation(nums=%v) returned: %v; expected: %v",
					test.nums, got, expected)
			}
		})
	}
}
