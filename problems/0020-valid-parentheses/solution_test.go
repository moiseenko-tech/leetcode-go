// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Andrei Moiseenko

package p0020

import "testing"

var tests = map[string]struct {
	s      string
	result bool
}{
	"()--true": {
		s:      "()",
		result: true,
	},
	"()[]{}--true": {
		s:      "()[]{}",
		result: true,
	},
	"(]--false": {
		s:      "(]",
		result: false,
	},
	"([])--true": {
		s:      "([])",
		result: true,
	},
	"([)]--false": {
		s:      "([)]",
		result: false,
	},
	"[--false": {
		s:      "[",
		result: false,
	},
}

func TestIsValid(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := isValid(test.s)
			expected := test.result
			if got != expected {
				t.Fatalf("isValid(s=%v) returned: %v; expected: %v",
					test.s, got, expected)
			}
		})
	}
}
