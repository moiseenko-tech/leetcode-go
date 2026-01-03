// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p0050

import (
	"math"
	"testing"
)

const eps = 1e-9

var tests = map[string]struct {
	x      float64
	n      int
	result float64
}{
	"2.0-10": {
		x:      2.0,
		n:      10,
		result: 1024.0,
	},
	"2.1-3": {
		x:      2.1,
		n:      3,
		result: 9.261,
	},
	"2.0--2": {
		x:      2.0,
		n:      -2,
		result: 0.25,
	},
}

func TestMyPow(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := myPow(test.x, test.n)
			expected := test.result
			if math.Abs(expected-got) > eps {
				t.Fatalf("myPow(x=%v, n=%v) returned: %v; expected: %v",
					test.x, test.n, got, expected)
			}
		})
	}
}
