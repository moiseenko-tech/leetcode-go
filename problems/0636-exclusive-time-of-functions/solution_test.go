// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p0636

import (
	"reflect"
	"testing"
)

var tests = map[string]struct {
	n      int
	logs   []string
	result []int
}{
	"2--0s0-1s2-1e5-0e6": {
		n:      2,
		logs:   []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"},
		result: []int{3, 4},
	},
	"1--0s0-0s2-0e5-0s6-0e6-0e7": {
		n:      1,
		logs:   []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"},
		result: []int{8},
	},
	"2--0s0-0s2-0e5-1s6-1e6-0e7": {
		n:      2,
		logs:   []string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"},
		result: []int{7, 1},
	},
}

func TestExclusiveTime(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := exclusiveTime(test.n, test.logs)
			expected := test.result
			if !reflect.DeepEqual(got, expected) {
				t.Fatalf("exclusiveTime(n=%v, logs=%v) returned: %v; expected: %v",
					test.n, test.logs, got, expected)
			}
		})
	}
}
