// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p1441

import (
	"reflect"
	"testing"
)

var tests = map[string]struct {
	target []int
	n      int
	result []string
}{
	"1-3--3": {
		target: []int{1, 3},
		n:      3,
		result: []string{"Push", "Push", "Pop", "Push"},
	},
	"1-2-3--3": {
		target: []int{1, 2, 3},
		n:      3,
		result: []string{"Push", "Push", "Push"},
	},
	"1-2--4": {
		target: []int{1, 2},
		n:      4,
		result: []string{"Push", "Push"},
	},
}

func TestBuildArray(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := buildArray(test.target, test.n)
			expected := test.result
			if !reflect.DeepEqual(got, expected) {
				t.Fatalf("buildArray(target=%v, n=%v) returned: %v; expected: %v",
					test.target, test.n, got, expected)
			}
		})
	}
}
